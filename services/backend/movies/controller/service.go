package movie

import (
	. "moovio/libs/businesslogic"
	"moovio/libs/helper"
	"moovio/libs/models"

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

func (s *MovieService) GetMovieByID(id string) (models.MovieModel, error) {
	var err error

	objid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.MovieModel{}, err
	}

	out, err := GetMovieByID(s.db, objid)
	if err != nil {
		return models.MovieModel{}, err
	}

	return out, nil
}
