package movie

import (
	"log"
	"moovio/libs/helper"
	"net/http"
)

type MoviesApiServer struct {
	svc MovieService
}

func NewMoviesApiServer(svc MovieService) *MoviesApiServer {
	return &MoviesApiServer{
		svc: svc,
	}
}

func (s *MoviesApiServer) Start(listenaddr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/getmovielist", s.GetMovieList)

	server := new(http.Server)
	server.Handler = mux
	server.Addr = listenaddr

	log.Println("Movie services running on", listenaddr)

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *MoviesApiServer) GetMovieList(w http.ResponseWriter, r *http.Request) {
	datas, err := s.svc.GetMovieList()

	if err != nil {
		helper.WriteJSON(w, http.StatusUnprocessableEntity, err.Error(), nil)
	}

	helper.WriteJSON(w, http.StatusOK, "", datas)
}
