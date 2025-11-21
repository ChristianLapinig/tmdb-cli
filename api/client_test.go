package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChristianLapinig/tmdb-cli/categories"
	"github.com/ChristianLapinig/tmdb-cli/models"
)

func TestClient_FetchMovies_Success(t *testing.T) {
	mockRes := models.MovieRes{
		Results: []models.Movie{
			{
				OriginalTitle: "Test Movie",
				Overview:      "Test Overview",
				Language:      "en",
				ReleaseDate:   "2024-01-01",
				Rating:        8.5,
			},
		},
	}

	// create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// verify request
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Error("Expected Authorization header")
		}
		if r.Header.Get("accept") != "application/json" {
			t.Error("Expected accept header to be application/json")
		}

		expectedPath := "/movie/now_playing"
		if r.URL.Path != expectedPath {
			t.Errorf("FAILED: expected path %s, got %s", expectedPath, r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockRes)
	}))
	defer server.Close()

	// create client
	client := CreateClient(server.URL, "test-token")
	result, err := client.FetchMovies(string(categories.NowPlaying))
	if err != nil {
		t.Fatalf("FAILED: error thrown %v", err)
	}
	if len(result.Results) == 0 {
		t.Error("FAILED: expected 1 movie result, got none.")
	}

	actualTitle := result.Results[0].OriginalTitle
	expectedTitle := mockRes.Results[0].OriginalTitle
	if actualTitle != expectedTitle {
		t.Errorf("FAILED: expected title %s, got %s", expectedTitle, actualTitle)
	}
}

func TestClient_FetchMovies_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	client := CreateClient(server.URL, "test-token")
	_, err := client.FetchMovies("now_playing")
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestClient_FetchMovies_EmptyResults(t *testing.T) {
	mockRes := models.MovieRes{
		Results: []models.Movie{},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockRes)
	}))
	defer server.Close()

	client := CreateClient(server.URL, "test-token")
	res, err := client.FetchMovies("now_playing")
	if err != nil {
		t.Fatal("ERROR:", err)
	}

	if len(res.Results) != 0 {
		t.Errorf("FAILED: Expected empty results, got %d", len(res.Results))
	}
}
