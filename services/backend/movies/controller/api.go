package movie

import (
	"log"
	"moovio/libs/helper"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/bson"
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

	mux.HandleFunc("/movies/getmovielist", s.GetMovieList)
	mux.HandleFunc("/movies/getmoviedetail", s.GetMovieDetail)
	mux.HandleFunc("/movies/getallmovies", s.GetAllMovies)

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
	var err error
	datas := []bson.M{}

	listtype := r.URL.Query().Get("list")

	if listtype == "" {
		helper.WriteJSON(w, http.StatusUnprocessableEntity, "List Type Not Defined", nil)
	}

	switch listtype {
	case "latest":
		datas, err = s.svc.GetLatestMovieList()
	case "imdb":
		datas, err = s.svc.GetTopImdbRating()
	}

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

func (s *MoviesApiServer) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	pageidstr := r.URL.Query().Get("page")
	pageid, _ := strconv.Atoi(pageidstr)

	data,err := s.svc.GetAllMovies(pageid)
	if err != nil {
		helper.WriteJSON(w, http.StatusUnprocessableEntity, err.Error(), nil)
	}

	helper.WriteJSON(w, http.StatusOK, "", data)
}