package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ChristianLapinig/tmdb-cli/constants"
	"github.com/ChristianLapinig/tmdb-cli/models"
)

type MovieClient interface {
	FetchMovies(category string) (*models.MovieRes, error)
}

var _ MovieClient = (*Client)(nil)

type Client struct {
	BaseURL     string
	HTTPClient  *http.Client
	AccessToken string
}

func CreateClient(baseURL, accessToken string) *Client {
	return &Client{
		BaseURL:     baseURL,
		AccessToken: accessToken,
		HTTPClient:  http.DefaultClient,
	}
}

func DefaultClient(accessToken string) *Client {
	return CreateClient(constants.BaseURL, accessToken)
}

func (c *Client) FetchMovies(category string) (*models.MovieRes, error) {
	url := fmt.Sprintf("%s/movie/%s", c.BaseURL, category)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %w", err)
	}

	var movieRes models.MovieRes
	if err := json.Unmarshal(body, &movieRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &movieRes, nil
}
