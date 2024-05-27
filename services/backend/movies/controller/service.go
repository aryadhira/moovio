package movie

import (
	. "moovio/libs/businesslogic"
	"moovio/libs/helper"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieService struct {
	db helper.Mongodbhelper
}

func NewMovieService(db helper.Mongodbhelper) MovieService {
	return MovieService{
		db: db,
	}
}

func (s *MovieService) GetLatestMovieList() ([]bson.M, error) {
	var err error
	out, err := GetLatestMovieList(s.db)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (s *MovieService) GetMovieByID(id string) (bson.M, error) {
	var err error

	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	out, err := GetMovieByID(s.db, objid)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (s *MovieService) GetTopImdbRating() ([]bson.M, error) {
	var err error
	out, err := GetTopImdbRating(s.db)
	if err != nil {
		return out, err
	}

	return out, nil
}

func (s *MovieService) GetAllMovies(page int) (bson.M, error) {
	var err error
	out, err := GetAllMovies(s.db, page)
	if err != nil {
		return out, err
	}	

	return out, nil
}