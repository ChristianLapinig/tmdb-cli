package program

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ChristianLapinig/tmdb-cli/categories"
	"github.com/ChristianLapinig/tmdb-cli/models"
)

type MockMovieClient struct {
	FetchMoviesFunc func(category string) (*models.MovieRes, error)
}

func (m *MockMovieClient) FetchMovies(category string) (*models.MovieRes, error) {
	if m.FetchMoviesFunc != nil {
		return m.FetchMoviesFunc(category)
	}
	return &models.MovieRes{}, nil
}

func TestCreateProgram(t *testing.T) {
	client := &MockMovieClient{}
	program := CreateProgram(client)
	if program.MovieClient != client {
		t.Error("FAILED: MovieClient was not set correctly")
	}
	if program.Output != os.Stdout {
		t.Error("FAILED: Output was not set correctly")
	}
}

func TestProgram_Run_Success(t *testing.T) {
	mockClient := &MockMovieClient{
		FetchMoviesFunc: func(category string) (*models.MovieRes, error) {
			return &models.MovieRes{
				Results: []models.Movie{
					{
						OriginalTitle: "Test Movie",
						Overview:      "Test Description",
						Language:      "en",
						ReleateDate:   "2024-01-01",
						Rating:        8.5,
					},
				},
			}, nil
		},
	}

	var buf bytes.Buffer
	program := CreateProgram(mockClient)
	program.Output = &buf

	// verify output is set to the buffer
	if program.Output != &buf {
		t.Fatal("Output was not set to buffer")
	}

	err := program.Run("now_playing")
	if err != nil {
		t.Fatal("ERROR:", err)
	}

	actual := buf.String()

	if actual == "" {
		t.Fatal("Buffer is empty - output is going somewhere else")
	}

	t.Logf("Buffer length: %d", buf.Len())
	t.Logf("Actual output: %s", actual)

	expectedOutput := []string{
		"Title: Test Movie",
		"Overview: Test Description",
		"Language: en",
		"Release Date: 2024-01-01",
		"Rating: 8.5",
	}

	for _, expected := range expectedOutput {
		if !strings.Contains(actual, expected) {
			t.Errorf("FAILED: expected output to contain %s, got %s", expected, actual)
		}
	}
}

func TestProgram_Run_InvalidCategory(t *testing.T) {
	client := &MockMovieClient{}
	program := CreateProgram(client)
	category := "invalid_category"
	err := program.Run(category)
	if err == nil {
		t.Fatal("ERROR: expected error, got nil")
	}

	expected := fmt.Sprintf("%s is not a valid category", category)
	if err.Error() != expected {
		t.Errorf("FAILED: expected error to be %s, got %s", expected, err.Error())
	}
}

func TestProgram_Run_FetchError(t *testing.T) {
	// mock error
	apiError := fmt.Errorf("Network timeout")
	client := &MockMovieClient{
		FetchMoviesFunc: func(category string) (*models.MovieRes, error) {
			return nil, apiError
		},
	}
	program := CreateProgram(client)
	err := program.Run("now_playing")

	if err == nil {
		t.Fatal("ERROR: expected error, got nil")
	}

	expected := fmt.Sprintf("failed to fetch movies: %s", apiError.Error())
	if err.Error() != expected {
		t.Errorf("FAILED: expected error to be %s, got %s", expected, err.Error())
	}
}

func TestProgram_Run_EmptyResults(t *testing.T) {
	client := &MockMovieClient{
		FetchMoviesFunc: func(category string) (*models.MovieRes, error) {
			return &models.MovieRes{
				Results: []models.Movie{},
			}, nil
		},
	}
	program := CreateProgram(client)
	err := program.Run("now_playing")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	expected := "No movies found for category now_playing"
	if err.Error() != expected {
		t.Errorf("FAILED: expected error to be %s, got %s", expected, err.Error())
	}
}

func TestProgram_Run_MultipleResults(t *testing.T) {
	client := &MockMovieClient{
		FetchMoviesFunc: func(category string) (*models.MovieRes, error) {
			return &models.MovieRes{
				Results: []models.Movie{
					{
						OriginalTitle: "Movie 1",
						Overview:      "Description 1",
						Language:      "en",
						ReleateDate:   "2024-01-01",
						Rating:        8.5,
					},
					{
						OriginalTitle: "Movie 2",
						Overview:      "Description 2",
						Language:      "fr",
						ReleateDate:   "2024-02-01",
						Rating:        7.2,
					},
				},
			}, nil
		},
	}

	var buf bytes.Buffer
	program := CreateProgram(client)
	program.Output = &buf
	err := program.Run("popular")
	if err != nil {
		t.Fatalf("ERROR: %v", err)
	}

	actual := buf.String()
	t.Logf("Actual output: %s", actual)

	if !strings.Contains(actual, "Movie 1") || !strings.Contains(actual, "Movie 2") {
		t.Error("FAILED: Not all movies are displayed.")
	}

	separatorCount := strings.Count(actual, "-----------------------")
	if separatorCount != 2 {
		t.Errorf("FAILED: Expected 2 separators, got %d", separatorCount)
	}
}

func TestProgram_Run_AllValidCategories(t *testing.T) {
	client := &MockMovieClient{
		FetchMoviesFunc: func(category string) (*models.MovieRes, error) {
			return &models.MovieRes{
				Results: []models.Movie{
					{
						OriginalTitle: "Movie 1",
						Overview:      "Description 1",
						Language:      "en",
						ReleateDate:   "2024-01-01",
						Rating:        8.5,
					},
				},
			}, nil
		},
	}

	for _, category := range categories.Categories {
		t.Run(string(category), func(t *testing.T) {
			var buf bytes.Buffer
			program := CreateProgram(client)
			program.Output = &buf
			err := program.Run(string(category))
			if err != nil {
				t.Errorf("FAILED: expected no error for category %s, got %v", category, err)
			}
		})
	}
}
