package helper

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbhelperInterface interface {
	FindOne(collection string, filter, option bson.M, result interface{}) error
	FindMany(collection string, filter, option bson.M, result interface{}) error
	Aggregate(collection string, pipeline bson.M, result interface{}) error
}

type Mongodbhelper struct {
	Client *mongo.Client
}

func InitDB() (Mongodbhelper, error) {
	uri := MongodbURIGenerator()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return Mongodbhelper{}, err
	}

	return Mongodbhelper{
		Client: client,
	}, nil
}

func (c *Mongodbhelper) FindOne(collection string, filter, option bson.M, result interface{}) error {
	dbname := os.Getenv("DB_NAME")
	coll := c.Client.Database(dbname).Collection(collection)

	err := coll.FindOne(context.TODO(), filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		return errors.New("no document found")
	}

	if err != nil {
		return err
	}

	return nil
}

func (c *Mongodbhelper) FindMany(collection string, filter, option bson.M, result interface{}) error {
	dbname := os.Getenv("DB_NAME")
	coll := c.Client.Database(dbname).Collection(collection)

	csr, err := coll.Find(context.TODO(), filter)
	if err == mongo.ErrNoDocuments {
		return errors.New("no document found")
	}

	if err != nil {
		return err
	}

	csr.All(context.TODO(), result)
	return nil
}

func (c *Mongodbhelper) Aggregate(collection string, pipeline []bson.M, result interface{}) error {
	dbname := os.Getenv("DB_NAME")
	coll := c.Client.Database(dbname).Collection(collection)

	csr, err := coll.Aggregate(context.TODO(), pipeline)
	if err == mongo.ErrNoDocuments {
		return errors.New("no document found")
	}

	if err != nil {
		return err
	}

	csr.All(context.TODO(), result)
	return nil
}

func (c *Mongodbhelper) InsertOne(collection string, data interface{}) error {
	dbname := os.Getenv("DB_NAME")
	coll := c.Client.Database(dbname).Collection(collection)

	_, err := coll.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Mongodbhelper) InsertMany(collection string, data []interface{}) error {
	dbname := os.Getenv("DB_NAME")
	coll := c.Client.Database(dbname).Collection(collection)

	_, err := coll.InsertMany(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Mongodbhelper) Delete(collection string, filter bson.M) error {
	dbname := os.Getenv("DB_NAME")
	coll := c.Client.Database(dbname).Collection(collection)

	_, err := coll.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
