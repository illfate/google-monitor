package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/illfate/google-monitor/internal/monitor"
)

type Mongo struct {
	collection *mongo.Collection
}

func NewMongo(collection *mongo.Collection) *Mongo {
	return &Mongo{
		collection: collection,
	}
}

func (m *Mongo) InsertRequestRes(ctx context.Context, res monitor.MonitorResult) error {
	_, err := m.collection.InsertOne(ctx, bson.M{"code": res.Code})
	if err != nil {
		return fmt.Errorf("failed to insert res: %w", err)
	}
	return nil
}
