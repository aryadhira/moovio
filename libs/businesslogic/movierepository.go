package businesslogic

import (
	"moovio/libs/constant"
	"moovio/libs/helper"
	"moovio/libs/models"
	"os"

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
		constant.Limit: 16,
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

func GetMovieByID(db helper.Mongodbhelper, id primitive.ObjectID) (bson.M, error) {
	var err error
	out := models.MovieModel{}

	err = db.FindOne(new(models.MovieModel).CollectionName(), bson.M{"_id": id}, bson.M{}, &out)
	if err != nil {
		return nil, err
	}

	title := out.Title
	moviemap := helper.StructToMap(out)
	
	matchstage := bson.M{
		constant.Match : bson.M{
			"title": title,
		},
	}

	groupstage := bson.M{
		constant.Group : bson.M{
			"_id" : "$quality",
		},
	}

	sortstage := bson.M{
		constant.Sort : bson.M{
			"_id" : 1,
		},
	}

	pipes := []bson.M{matchstage,groupstage,sortstage}

	resquality := []bson.M{}
	err = db.Aggregate(new(models.MovieModel).CollectionName(),pipes,&resquality)
	if err != nil {
		return nil, err
	}

	qualities := []string{}
	for _,each := range resquality {
		qualities = append(qualities, helper.InterfaceToString(each["_id"]))
	}

	moviemap["qualities"] = qualities

	return moviemap, nil
}

func GetTopImdbRating(db helper.Mongodbhelper) ([]bson.M, error) {
	var err error
	out := []bson.M{}

	groupstage := bson.M{
		constant.Group: bson.M{
			"_id":  "$title",
			"data": bson.M{"$first": "$$ROOT"},
		},
	}

	projectstage := bson.M{
		constant.Project: bson.M{
			"_id":    0,
			"id":     "$data._id",
			"title":  "$_id",
			"rating": "$data.rating",
			"year":   "$data.year",
			"cover":  "$data.cover",
		},
	}

	sortstage := bson.M{
		constant.Sort: bson.M{"rating": -1},
	}

	limitstage := bson.M{
		constant.Limit: 16,
	}

	pipe := []bson.M{}
	pipe = append(pipe, groupstage)
	pipe = append(pipe, projectstage)
	pipe = append(pipe, sortstage)
	pipe = append(pipe, limitstage)

	err = db.Aggregate(new(models.MovieModel).CollectionName(), pipe, &out)
	if err != nil {
		return out, err
	}

	return out, nil
}

func GetAllMovies(db helper.Mongodbhelper, index int) (bson.M, error) {
	var err error

	groupstage := bson.M{
		constant.Group: bson.M{
			"_id":  "$title",
			"data": bson.M{"$first": "$$ROOT"},
		},
	}

	projectstage := bson.M{
		constant.Project: bson.M{
			"_id":    0,
			"id":     "$data._id",
			"title":  "$_id",
			"rating": "$data.rating",
			"year":   "$data.year",
			"cover":  "$data.cover",
		},
	}

	pipe := []bson.M{groupstage,projectstage}
	movies := []bson.M{}
	err = db.Aggregate(new(models.MovieModel).CollectionName(), pipe, &movies)
	if err != nil {
		return nil, err
	}

	totallist := len(movies)
	movieperpages := helper.InterfaceToInt(os.Getenv("MOVIE_PER_PAGE"))
	totalpages := totallist / movieperpages

	skip := (index - 1) * movieperpages

	skipstage := bson.M{constant.Skip:skip}
	limitstage := bson.M{constant.Limit:movieperpages}

	pipe = append(pipe, skipstage)
	pipe = append(pipe, limitstage)

	limitedmovies := []bson.M{}
	err = db.Aggregate(new(models.MovieModel).CollectionName(), pipe, &limitedmovies)
	if err != nil {
		return nil, err
	}

	result := bson.M{
		"totalmovies" : totallist,
		"movieperpages" : movieperpages,
		"totalpages": totalpages,
		"movies": limitedmovies,
	}


	return result, nil
}