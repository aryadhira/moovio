package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieModel struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	Title           string             `json:"title" bson:"title"`
	Year            int                `json:"year" bson:"year"`
	Synopsis        string             `json:"synopsis" bson:"synopsis"`
	Rating          float64            `json:"rating" bson:"rating"`
	Cover           string             `json:"cover" bson:"cover"`
	Quality         string             `json:"quality" bson:"quality"`
	Hash            string             `json:"hash" bson:"hash"`
	Slug            string             `json:"slug" bson:"slug"`
	MagnetUrl       string             `json:"magneturl" bson:"magneturl"`
	Category        []string           `json:"category" bson:"category"`
	PopulateDate    time.Time          `json:"populatedate" bson:"populatedate"`
	PopulateDateInt int                `json:"populatedateint" bson:"populatedateint"`
}

func (m *MovieModel) NewMovieModel() *MovieModel {
	return &MovieModel{
		Id: primitive.NewObjectID(),
	}
}

func (m *MovieModel) CollectionName() string {
	return "Movies"
}
