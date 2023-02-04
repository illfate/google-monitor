//go:build real_api

package googleclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRealClientMakeGetRequest(t *testing.T) {
	c := NewClient(http.DefaultClient, "https://google.com")
	got, err := c.MakeGetRequest(context.Background())
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, got.Code)
}
