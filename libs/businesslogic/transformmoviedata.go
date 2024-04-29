package businesslogic

import (
	"errors"
	"log"
	"moovio/libs/helper"
	"moovio/libs/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TransformMovieData(db helper.Mongodbhelper, datas map[string]interface{}) error {
	log.Println("Start Transforming Movie Data...")
	populatedate := time.Now()
	dataresult := datas["data"]
	if dataresult == nil {
		return errors.New("empty data result")
	}

	dataobj := datas["data"].(map[string]interface{})
	movies := dataobj["movies"]

	if movies == nil {
		return errors.New("empty movies data")
	}

	moviedatas := movies.([]interface{})
	movs := make([]interface{}, 0)

	log.Println("Getting existing movie data")
	// Current Data
	existingdata := []models.MovieModel{}
	err := db.FindMany(new(models.MovieModel).CollectionName(), bson.M{}, bson.M{}, &existingdata)
	if err != nil {
		return err
	}
	existingdatamap := make(map[string]bool)
	for _, each := range existingdata {
		key := each.Title + "|" + each.Hash
		existingdatamap[key] = true
	}

	log.Println("Transformation start...")
	moviecount := len(moviedatas)
	log.Printf("Total Movie to Transform %d", moviecount)
	for _, each := range moviedatas {
		obj := each.(map[string]interface{})

		movie := models.MovieModel{}
		movie.Title = helper.InterfaceToString(obj["title"])
		movie.Year = int(helper.InterfaceToFloat64(obj["year"]))
		movie.Cover = helper.InterfaceToString(obj["large_cover_image"])
		movie.Slug = helper.InterfaceToString(obj["slug"])
		movie.Rating = helper.InterfaceToFloat64(obj["rating"])
		movie.Synopsis = helper.InterfaceToString(obj["synopsis"])
		movie.PopulateDate = populatedate
		movie.PopulateDateInt = helper.ConvertDateTimetoDateInt(populatedate)
		if obj["genres"] != nil {
			movie.Category = helper.ArrayinterfaceToArrayString(obj["genres"].([]interface{}))
		}

		torrents := obj["torrents"].([]interface{})
		for _, torrent := range torrents {
			torrentobj := torrent.(map[string]interface{})
			movie.Id = primitive.NewObjectID()
			movie.Quality = helper.InterfaceToString(torrentobj["quality"])
			movie.Hash = helper.InterfaceToString(torrentobj["hash"])
			magneturl := helper.GenerateMagnetUrl(movie.Title, movie.Hash, movie.Quality)
			movie.MagnetUrl = magneturl

			// Duplication Validation
			if !DuplicateCheck(movie, existingdatamap) {
				movs = append(movs, movie)
			}
		}
	}
	log.Println(len(movs))
	log.Println("Start Inserting Movie Data...")
	err = db.InsertMany(new(models.MovieModel).CollectionName(), movs)
	if err != nil {
		return err
	}

	log.Println("Transform Movie Data Done at:", time.Since(populatedate))

	return nil
}

func DuplicateCheck(movie models.MovieModel, existingmap map[string]bool) bool {
	out := false

	key := movie.Title + "|" + movie.Hash
	out = existingmap[key]

	return out
}
