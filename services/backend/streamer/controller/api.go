package streamer

import (
	"log"
	"moovio/libs/helper"
	"net/http"
	"sync"

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

	mux.HandleFunc("/stream", s.HandleStream)
	mux.HandleFunc("/stopstream", s.StopStreaming)

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

func (s *StreamerApiServer) HandleStream(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	title := r.URL.Query().Get("title")
	quality := r.URL.Query().Get("quality")

	magnet, err := s.svc.GetMovieMagnetUrl(title, quality)
	if err != nil {
		helper.WriteJSON(w, http.StatusUnprocessableEntity, err.Error(), nil)
	}

	if magnet == "" {
		helper.WriteJSON(w, http.StatusInternalServerError, "Magnet URL are not available", nil)
		return
	}

	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "./torrent_data"
	client, err := torrent.NewClient(cfg)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, "Failed to start client", nil)
	}

	torrentDownload, err = client.AddMagnet(magnet)
	if err != nil {
		helper.WriteJSON(w, http.StatusInternalServerError, "Failed to start client", nil)
	}

	<-torrentDownload.GotInfo()

	for _, f := range torrentDownload.Files() {
		if file == nil || f.Length() > file.Length() {
			file = f
		}
	}

	if file == nil {
		helper.WriteJSON(w, http.StatusInternalServerError, "Torrent File Not Found", nil)
		return
	}

	// Stream the file
	reader := file.NewReader()
	defer reader.Close()
	w.Header().Set("Content-Type", "video/mp4") // Adjust the content type accordingly
	// http.ServeContent(w, r, targetFile.DisplayPath(), time.Time{}, reader)
	bytesWritten := 0

	// Create a buffer to read from the file reader
	buffer := make([]byte, 32*1024)
	for {
		select {
		case <-r.Context().Done():
			// Client has closed the connection or signaled to stop streaming
			log.Println("Streaming stopped by client")
			return
		default:
			n, err := reader.Read(buffer)
			if err != nil && err.Error() != "EOF" {
				// Handle error
				log.Printf("Error reading file: %v\n", err)
				return
			}

			if n > 0 {
				// Write the buffer to the ResponseWriter
				_, err = w.Write(buffer[:n])
				if err != nil {
					// Handle error
					log.Printf("Error writing response: %v\n", err)
					return
				}

				// Increment bytes written
				bytesWritten += n

				// Log progress

				log.Printf("Progress: %d/%d bytes streamed\n", bytesWritten, file.Length())
			}
		}
	}

}

func (s *StreamerApiServer) StopStreaming(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	if torrentDownload != nil {
		// Remove the torrent download from the client
		torrentDownload.Drop()
	}

	if client != nil {
		// Close the torrent client
		client.Close()
		client = nil
	}

	log.Println("Streaming stopped")
}
