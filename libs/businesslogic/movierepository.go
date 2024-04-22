package businesslogic

import (
	"moovio/libs/constant"
	"moovio/libs/helper"
	"moovio/libs/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLatestMovieList(db helper.Mongodbhelper) ([]bson.M, error) {
	var err error
	out := []bson.M{}

	groupstage := bson.M{
		constant.Group: bson.M{
			"_id":  "$title",
			"data": bson.M{constant.First: constant.ROOT},
		},
	}

	projectstage := bson.M{
		constant.Project: bson.M{
			"_id":          0,
			"id":           "$data._id",
			"title":        "$_id",
			"cover":        "$data.cover",
			"rating":       "$data.rating",
			"year":         "$data.year",
			"populatedate": "$data.populatedate",
		},
	}

	sortstage := bson.M{
		constant.Sort: bson.M{
			"populatedate": -1,
			"title":        1,
		},
	}

	limitstage := bson.M{
		constant.Limit: 10,
	}

	pipes := []bson.M{}
	pipes = append(pipes, groupstage)
	pipes = append(pipes, projectstage)
	pipes = append(pipes, sortstage)
	pipes = append(pipes, limitstage)

	err = db.Aggregate(new(models.MovieModel).CollectionName(), pipes, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func GetMovieByID(db helper.Mongodbhelper, id primitive.ObjectID) (models.MovieModel, error) {
	var err error
	out := models.MovieModel{}

	err = db.FindOne(new(models.MovieModel).CollectionName(), bson.M{"_id": id}, bson.M{}, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}
