package movie

import (
	"log"
	"moovio/libs/helper"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	mux := mux.NewRouter()

	mux.HandleFunc("/getmovielist", s.GetMovieList)
	mux.HandleFunc("/getmoviedetail", s.GetMovieDetail)

	handler := cors.Default().Handler(mux)

	server := new(http.Server)
	server.Handler = handler
	server.Addr = listenaddr

	log.Println("Movie services running on", listenaddr)

	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *MoviesApiServer) GetMovieList(w http.ResponseWriter, r *http.Request) {
	datas, err := s.svc.GetLatestMovieList()

	if err != nil {
		helper.WriteJSON(w, http.StatusUnprocessableEntity, err.Error(), nil)
	}

	helper.WriteJSON(w, http.StatusOK, "", datas)
}

func (s *MoviesApiServer) GetMovieDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	data, err := s.svc.GetMovieByID(id)

	if err != nil {
		helper.WriteJSON(w, http.StatusUnprocessableEntity, err.Error(), nil)
	}

	helper.WriteJSON(w, http.StatusOK, "", data)
}
