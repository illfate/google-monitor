package monitor_test

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/illfate/google-monitor/internal/mock"
	"github.com/illfate/google-monitor/internal/monitor"
)

func TestServiceMonitor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("happy path", func(t *testing.T) {
		repo := mock.NewMockRepository(ctrl)
		client := mock.NewMockGoogleClient(ctrl)

		repo.EXPECT().InsertRequestRes(gomock.Any(), monitor.RequestResult{Code: 200}).Return(nil)
		client.EXPECT().MakeGetRequest(gomock.Any()).Return(monitor.RequestResult{Code: 200}, nil)

		service := monitor.NewService(client, repo)
		err := service.Monitor(context.Background())
		assert.NoError(t, err)
	})
	t.Run("client returns error", func(t *testing.T) {
		repo := mock.NewMockRepository(ctrl)
		client := mock.NewMockGoogleClient(ctrl)

		client.EXPECT().MakeGetRequest(gomock.Any()).Return(monitor.RequestResult{}, io.EOF)

		service := monitor.NewService(client, repo)
		err := service.Monitor(context.Background())
		require.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("failed to make get request: %w", io.EOF), err)
	})
	t.Run("repo return error", func(t *testing.T) {
		repo := mock.NewMockRepository(ctrl)
		client := mock.NewMockGoogleClient(ctrl)

		client.EXPECT().MakeGetRequest(gomock.Any()).Return(monitor.RequestResult{Code: 200}, nil)
		repo.EXPECT().InsertRequestRes(gomock.Any(), monitor.RequestResult{Code: 200}).Return(io.EOF)

		service := monitor.NewService(client, repo)
		err := service.Monitor(context.Background())
		require.NotNil(t, err)
		assert.Equal(t, fmt.Errorf("failed to insert request result: %w", io.EOF), err)
	})
}
