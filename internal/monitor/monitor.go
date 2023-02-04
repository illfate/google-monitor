// Package monitor implement monitoring logic.
package monitor

import (
	"context"
	"fmt"
)

type GoogleClient interface {
	MakeGetRequest(ctx context.Context) (RequestResult, error)
}

type Repository interface {
	InsertRequestRes(ctx context.Context, res RequestResult) error
}

type RequestResult struct {
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

// Monitor makes request to client and store the API result.
// It assumes we dont wont to store error.
func (s *Service) Monitor(ctx context.Context) error {
	res, err := s.client.MakeGetRequest(ctx)
	if err != nil {
		return fmt.Errorf("failed to make get request: %w", err)
	}
	err = s.repo.InsertRequestRes(ctx, res)
	if err != nil {
		return fmt.Errorf("failed to insert request result: %w", err)
	}
	return nil
}
