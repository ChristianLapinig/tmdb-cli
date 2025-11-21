package program

import (
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/ChristianLapinig/tmdb-cli/api"
	"github.com/ChristianLapinig/tmdb-cli/categories"
	"github.com/ChristianLapinig/tmdb-cli/models"
)

type Program struct {
	MovieClient api.MovieClient
	Output      io.Writer
}

func CreateProgram(client api.MovieClient) *Program {
	return &Program{
		MovieClient: client,
		Output:      os.Stdout,
	}
}

func (p *Program) Run(category string) error {
	if !slices.Contains(categories.Categories, categories.Category(category)) {
		return fmt.Errorf("%s is not a valid category", category)
	}

	movieRes, err := p.MovieClient.FetchMovies(category)
	if err != nil {
		return fmt.Errorf("failed to fetch movies: %w", err)
	}

	if len(movieRes.Results) == 0 {
		return fmt.Errorf("No movies found for category %s", category)
	}

	p.displayMovies(movieRes.Results)
	return nil
}

// Legacy function for backwards compatability
func Execute(category, accessToken string) {
	client := api.DefaultClient(accessToken)
	program := CreateProgram(client)
	if err := program.Run(category); err != nil {
		fmt.Fprint(os.Stderr, "Error:", err)
	}
}

func (p *Program) displayMovies(movies []models.Movie) {
	for _, movie := range movies {
		fmt.Fprintln(p.Output, "-----------------------")
		fmt.Fprintf(p.Output, "Title: %s\n", movie.OriginalTitle)
		fmt.Fprintf(p.Output, "Overview: %s\n", movie.Overview)
		fmt.Fprintf(p.Output, "Language: %s\n", movie.Language)
		fmt.Fprintf(p.Output, "Release Date: %s\n", movie.ReleateDate)
		fmt.Fprintf(p.Output, "Rating: %.1f\n", movie.Rating)
	}
}
