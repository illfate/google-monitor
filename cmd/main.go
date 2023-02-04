package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/illfate/google-monitor/internal/googleclient"
	"github.com/illfate/google-monitor/internal/monitor"
	"github.com/illfate/google-monitor/internal/repository"
)

type Config struct {
	MongoURI        string `envconfig:"MONGO_URI" default:"mongodb://root:example@localhost:27017/?ssl=false&authSource=admin"`
	MongoCollection string `envconfig:"MONGO_COLLECTION" default:"response_results"`
	MongoDatabase   string `envconfig:"MONGO_DATABASE" default:"monitor"`
	GoogleURL       string `envconfig:"GOOGLE_URL" default:"https://google.com"`
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return fmt.Errorf("failed to prcoess envs: %w", err)
	}
	client, err := connectToMongo(cfg.MongoURI)
	if err != nil {
		return fmt.Errorf("failed to connect to mongo: %w", err)
	}
	defer func() {
		ctx, cancelF := context.WithTimeout(context.Background(), time.Second)
		defer cancelF()
		err := client.Disconnect(ctx)
		if err != nil {
			log.Printf("Failed to disconnect from mongo: %s", err)
		}
	}()

	collection := client.Database(cfg.MongoDatabase).Collection(cfg.MongoCollection)
	repo := repository.NewMongo(collection)
	httpClient := &http.Client{
		Timeout: 3 * time.Second,
	}

	googleClient := googleclient.NewClient(httpClient, cfg.GoogleURL)

	service := monitor.NewService(googleClient, repo)
	err = runMonitor(service)
	if err != nil {
		return err
	}
	return nil
}

func runMonitor(service *monitor.Service) error {
	ctx, cancelF := context.WithTimeout(context.Background(), 2*time.Second) // TODO: move to cfg
	defer cancelF()

	res, err := service.Monitor(ctx)
	if err != nil {
		return fmt.Errorf("failed to run monitor: %w", err)
	}
	msg := fmt.Sprintf("code is %d\n", res.Code)
	_, err = io.Copy(os.Stdout, strings.NewReader(msg))
	if err != nil {
		return fmt.Errorf("failed to output result: %w", err)
	}
	return nil
}

func connectToMongo(uri string) (*mongo.Client, error) {
	ctx, cancelF := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelF()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
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
