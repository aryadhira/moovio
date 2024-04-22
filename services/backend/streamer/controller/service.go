package streamer

import (
	"moovio/libs/helper"
	"moovio/libs/models"

	"go.mongodb.org/mongo-driver/bson"
)

type StreamerService struct {
	db helper.Mongodbhelper
}

func NewStreamerService(db helper.Mongodbhelper) StreamerService {
	return StreamerService{
		db: db,
	}
}

func (s *StreamerService) GetMovieMagnetUrl(title, quality string) (string, error) {
	out := ""

	movie := models.MovieModel{}
	err := s.db.FindOne(new(models.MovieModel).CollectionName(), bson.M{"title": title, "quality": quality}, bson.M{}, &movie)
	if err != nil {
		return out, err
	}

	out = movie.MagnetUrl

	return out, nil
}
