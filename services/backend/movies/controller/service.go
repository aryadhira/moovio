package movie

import (
	"moovio/libs/constant"
	"moovio/libs/helper"
	"moovio/libs/models"

	"go.mongodb.org/mongo-driver/bson"
)

type MovieService struct {
	db helper.Mongodbhelper
}

func NewMovieService(db helper.Mongodbhelper) MovieService {
	return MovieService{
		db: db,
	}
}

func (s *MovieService) GetMovieList() ([]bson.M, error) {
	var err error
	out := []bson.M{}

	groupstage := bson.M{
		constant.Group: bson.M{
			"_id": bson.M{
				"title":  "$title",
				"cover":  "$cover",
				"rating": "$rating",
			},
		},
	}

	projectstage := bson.M{
		constant.Project: bson.M{
			"_id":    0,
			"title":  "$_id.title",
			"cover":  "$_id.cover",
			"rating": "$_id.rating",
		},
	}

	sortstage := bson.M{
		constant.Sort: bson.M{
			"populatedate": 1,
		},
	}

	pipes := []bson.M{}
	pipes = append(pipes, sortstage)
	pipes = append(pipes, groupstage)
	pipes = append(pipes, projectstage)

	err = s.db.Aggregate(new(models.MovieModel).CollectionName(), pipes, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}
