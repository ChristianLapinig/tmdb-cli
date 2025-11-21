package api

import (
	"github.com/ChristianLapinig/tmdb-cli/constants"
	"github.com/ChristianLapinig/tmdb-cli/models"
)

// FetchMovies is a convenience function that uses the default client
func FetchMovies(category, accessToken string) (*models.MovieRes, error) {
	client := CreateClient(constants.BaseURL, accessToken)
	return client.FetchMovies(category)
}

// DefaultClient that uses default settings
func DefaultClient(accessToken string) *Client {
	return CreateClient(constants.BaseURL, accessToken)
}
