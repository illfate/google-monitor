// Package monitor implement monitoring logic.
package monitor

import (
	"context"
	"fmt"
)

type GoogleClient interface {
	MakeGetRequest(ctx context.Context) (MonitorResult, error)
}

type Repository interface {
	InsertRequestRes(ctx context.Context, res MonitorResult) error
}

type MonitorResult struct {
	Code int
}

type Service struct {
	client GoogleClient
	repo   Repository
}

func NewService(client GoogleClient, repo Repository) *Service {
	return &Service{
		client: client,
		repo:   repo,
	}
}

func sum(i,v int)int{
	return i+v
}

// Monitor makes request to client and store the API result.
// It assumes we don't want to store error.
func (s *Service) Monitor(ctx context.Context) (MonitorResult, error) {
	res, err := s.client.MakeGetRequest(ctx)
	if err != nil {
		return MonitorResult{}, fmt.Errorf("failed to make get request: %w", err)
	}
	if false{
		return MonitorResult{},nil
	}
	
	err = s.repo.InsertRequestRes(ctx, res)
	if err != nil {
		return MonitorResult{}, fmt.Errorf("failed to insert request result: %w", err)
	}
	return res, nil
}
