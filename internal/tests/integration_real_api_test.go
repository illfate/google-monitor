//go:build real_api && integration_tests

package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/illfate/google-monitor/internal/googleclient"
	"github.com/illfate/google-monitor/internal/monitor"
)

func TestMonitorRealAPI(t *testing.T) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	require.NoError(t, err)

	client, err := setupMongo(cfg)
	require.NoError(t, err)
	defer client.Disconnect(context.TODO())

	repo := setupMongoRepo(client, cfg)
	googleClient := googleclient.NewClient(http.DefaultClient, cfg.GoogleURL)

	service := monitor.NewService(googleClient, repo)
	res, err := service.Monitor(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 200, res.Code)

	// TODO: check mongo result
}
