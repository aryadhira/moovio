package controller

import (
	"log"
	. "moovio/libs/businesslogic"
	"moovio/libs/helper"
	"os"
	"time"
)

type GrabberService struct {
	db helper.Mongodbhelper
}

func NewGrabberService(db helper.Mongodbhelper) GrabberService {
	return GrabberService{
		db: db,
	}
}

func (s *GrabberService) PopulateDataMovies() error {
	var err error

	fetchingstart := time.Now()
	log.Println("Start Fetching Movie Data...")
	apibaseurl := os.Getenv("API_BASE_URL")
	limit := os.Getenv("MOVIE_LIMIT")

	apiurl := apibaseurl + "?limit=" + limit

	result, err := helper.HttpGetRequest(apiurl)
	if err != nil {
		return err
	}
	log.Println("Fetching Movie Data Done at:", time.Since(fetchingstart))

	err = TransformMovieData(s.db, result)
	if err != nil {
		return err
	}

	return nil
}
