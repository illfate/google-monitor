package googleclient

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientMakeGetRequest(t *testing.T) {
	const url = "https://google.com"

	t.Run("ok code", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", url,
			httpmock.NewStringResponder(200, `{}`))

		c := NewClient(http.DefaultClient, url)
		got, err := c.MakeGetRequest(context.Background())
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, got.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", url,
			httpmock.NewStringResponder(500, `{}`))

		c := NewClient(http.DefaultClient, url)
		got, err := c.MakeGetRequest(context.Background())
		require.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, got.Code)
	})
}
