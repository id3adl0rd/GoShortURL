package db

import (
	"context"
	"go-short-me/internal/infrastructure/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

// TODO Перенести в .env
const (
	timeout           = 100 * time.Millisecond
	dbName     string = "yetshorter"
	collection string = "urls"
)

type Mongo struct {
	Client *mongo.Client
}

func (m *Mongo) InsertURL(ctx context.Context, obj interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := getConn(m.Client).InsertOne(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) GetByURL(ctx context.Context, value string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	filter := bson.D{{"url", value}}
	var url models.UrlDB
	err := getConn(m.Client).FindOne(ctx, filter).Decode(&url)
	if err != nil {
		return "", err
	}

	return url.CreatedURL, nil
}

func (m *Mongo) GetAll(ctx context.Context) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result := make(map[string]string)
	filter := bson.D{}
	cursor, err := getConn(m.Client).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getConn(m *mongo.Client) *mongo.Collection {
	return m.Database(dbName).Collection(collection)
}
