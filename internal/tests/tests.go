// Package tests contains integration tests.
package tests

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/illfate/google-monitor/internal/repository"
)

type Config struct {
	MongoURI        string `envconfig:"MONGO_URI" default:"mongodb://root:example@localhost:27017/?ssl=false&authSource=admin"`
	MongoCollection string `envconfig:"MONGO_COLLECTION" default:"response_results"`
	MongoDatabase   string `envconfig:"MONGO_DATABASE" default:"monitor"`
	GoogleURL       string `envconfig:"GOOGLE_URL" default:"https://google.com"`
}

func setupMongoRepo(client *mongo.Client, cfg Config) *repository.Mongo {
	collection := client.Database(cfg.MongoDatabase).Collection(cfg.MongoCollection)
	repo := repository.NewMongo(collection)
	return repo
}

func setupMongo(cfg Config) (*mongo.Client, error) {
	ctx, cancelF := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelF()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return nil, err
	}

	pingCtx, pingCancelF := context.WithTimeout(context.Background(), 2*time.Second)
	defer pingCancelF()
	err = client.Ping(pingCtx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return client, nil
}
