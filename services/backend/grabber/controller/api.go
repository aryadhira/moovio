package controller

import (
	"log"
	"moovio/libs/helper"
	"net/http"
)

type GrabberApiServer struct {
	svc GrabberService
}

func NewGrabberApiServer(svc GrabberService) *GrabberApiServer {
	return &GrabberApiServer{
		svc: svc,
	}
}

func (s *GrabberApiServer) Start(listenaddr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/grabber/fetchmoviedata", s.handleFetchMovieData)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = listenaddr

	log.Println("Grabber services running on", listenaddr)

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *GrabberApiServer) handleFetchMovieData(w http.ResponseWriter, r *http.Request) {
	err := s.svc.PopulateDataMovies()
	if err != nil {
		helper.WriteJSON(w, http.StatusUnprocessableEntity, err.Error(), nil)
	}

	helper.WriteJSON(w, http.StatusOK, "", nil)
}
