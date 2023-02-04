// Package googleclient contains google client related logic.
package googleclient

import (
	"context"
	"fmt"
	"net/http"

	"github.com/illfate/google-monitor/internal/monitor"
)

// Client holds logic of accessing to google api.
type Client struct {
	client  *http.Client
	baseURl string
}

// NewClient builds new client.
func NewClient(client *http.Client, baseURL string) *Client {
	return &Client{
		client:  client,
		baseURl: baseURL,
	}
}

// MakeGetRequest makes get request to passed google url and returns status code.
func (c *Client) MakeGetRequest(ctx context.Context) (monitor.RequestResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURl, http.NoBody)
	if err != nil {
		return monitor.RequestResult{}, fmt.Errorf("failed to create a request: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return monitor.RequestResult{}, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	reqRes := monitor.RequestResult{
		Code: resp.StatusCode,
	}
	return reqRes, nil
}
