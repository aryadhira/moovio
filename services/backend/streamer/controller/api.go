package streamer

import (
	"fmt"
	"io"
	"log"
	"moovio/libs/helper"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/anacrolix/torrent"
)

var (
	client          *torrent.Client
	torrentDownload *torrent.Torrent
	file            *torrent.File
	mutex           sync.Mutex
)

type StreamerApiServer struct {
	svc StreamerService
}

func NewStreamerApiServer(svc StreamerService) *StreamerApiServer {
	return &StreamerApiServer{
		svc: svc,
	}
}

func (s *StreamerApiServer) Start(listenaddr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/streamer/stream", s.HandleStream)
	mux.HandleFunc("/streamer/localstream", s.LocalStream)
	mux.HandleFunc("/streamer/stopstream", s.StopStreaming)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = listenaddr

	log.Println("Streamer services running on", listenaddr)

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

// HandleStream handler to stream torrent
func (s *StreamerApiServer) HandleStream(w http.ResponseWriter, r *http.Request) {
	// Lock mutex to ensure thread safety
	mutex.Lock()
	defer mutex.Unlock()

	// Extract title and quality from URL query parameters
	title := r.URL.Query().Get("title")
	quality := r.URL.Query().Get("quality")

	// Retrieve magnet URL for the movie
	magnet, err := s.svc.GetMovieMagnetUrl(title, quality)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate magnet URL
	if magnet == "" {
		http.Error(w, "Magnet URL not available", http.StatusInternalServerError)
		return
	}

	// Configure torrent client
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "./torrent_data"
	client, err := torrent.NewClient(cfg)
	if err != nil {
		http.Error(w, "Failed to start client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Add magnet to client
	torrentDownload, err := client.AddMagnet(magnet)
	if err != nil {
		http.Error(w, "Failed to start download", http.StatusInternalServerError)
		return
	}

	// Wait for torrent to be ready
	<-torrentDownload.GotInfo()

	// Select largest file in torrent
	var targetFile *torrent.File
	for _, f := range torrentDownload.Files() {
		if targetFile == nil || f.Length() > targetFile.Length() {
			targetFile = f
		}
	}

	// Check if file is found
	if targetFile == nil {
		http.Error(w, "Torrent file not found", http.StatusInternalServerError)
		return
	}

	// Stream the file
	reader := targetFile.NewReader()
	defer reader.Close()

	// Set Content-Type header
	w.Header().Set("Content-Type", "video/mp4")

	// Serve the content
	http.ServeContent(w, r, targetFile.DisplayPath(), time.Time{}, reader)
}

// StopStreaming handler to stop stream and torrent download
func (s *StreamerApiServer) StopStreaming(w http.ResponseWriter, r *http.Request) {
	// Lock mutex to ensure thread safety
	mutex.Lock()
	defer mutex.Unlock()

	// Close torrent client if it's not nil
	if client != nil {
		client.Close()
		client = nil
	}

	// Drop torrent download if it's not nil
	if torrentDownload != nil {
		torrentDownload.Drop()
	}

	log.Println("Streaming stopped")
	// Respond with success message
	w.WriteHeader(http.StatusOK)
}

func (s *StreamerApiServer) LocalStream(w http.ResponseWriter, r *http.Request) {
	torrentdir, _ := os.ReadDir("./torrent_data")
	for _, items := range torrentdir {
		if items.IsDir() {
			log.Println(items.Name())
			moviefolder := filepath.Join("./torrent_data", items.Name())
			nesteddata, _ := os.ReadDir(moviefolder)
			for _, nestitem := range nesteddata {
				log.Println(filepath.Join(moviefolder, nestitem.Name()))
			}
		}
	}

	filePath := "./torrent_data/Downtown Owl (2023) [720p] [WEBRip] [YTS.MX]/Downtown.Owl.2023.720p.WEBRip.x264.AAC-[YTS.MX].mp4"

	video, err := os.Open(filePath)
	if err != nil {
		helper.WriteJSON(w, http.StatusNotFound, "File Not Found", nil)
		return
	}
	defer video.Close()

	stat, err := video.Stat()
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, "Error Reading File", nil)
		return
	}

	fileSize := stat.Size()
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))

	rangeHeader := r.Header.Get("Range")
	if rangeHeader == "" || rangeHeader == "bytes=0-" {
		io.Copy(w, video)
		return
	}

	var start, end int64
	_, err = fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)
	if err != nil || start < 0 || start >= fileSize || end < start {
		helper.WriteJSON(w, http.StatusInternalServerError, "Invalid Range", nil)
		return
	}

	if end == 0 || end >= fileSize {
		end = fileSize - 1
	}

	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
	w.Header().Set("Accept-Ranges", "bytes")
	w.WriteHeader(http.StatusPartialContent)

	video.Seek(start, 0)
	io.CopyN(w, video, end-start+1)
}
