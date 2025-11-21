package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ChristianLapinig/tmdb-cli/api"
	"github.com/ChristianLapinig/tmdb-cli/categories"
	"github.com/ChristianLapinig/tmdb-cli/constants"
	"github.com/ChristianLapinig/tmdb-cli/program"
)

func main() {
	category := flag.String("type", "", "Category of movies to fetch(now_playing, popular, upcoming, and top_rated)")

	flag.Usage = func() {
		fmt.Println("Usage: tmdb-cli --type <category>")
		fmt.Println("Please note that the --type flag must also be passed.")
		fmt.Printf("One of the following categories must be passed as an argument:\n")
		for _, category := range categories.Categories {
			fmt.Println(strings.Repeat(" ", 4) + "- " + string(category))
		}
		fmt.Println("You also need to ensure that you have a valid TMDB Read Access Token.")
		fmt.Println("This requires an account on The Movie Database (https://www.themoviedb.org/).")
	}

	accessToken := os.Getenv("TMDB_READ_ACCESS_TOKEN")
	if accessToken == "" {
		fmt.Println(constants.AccessTokenRequired)
		os.Exit(1)
	}

	flag.Parse()

	if *category == "" {
		flag.Usage()
		os.Exit(1)
	}

	client := api.DefaultClient(accessToken)
	program := program.CreateProgram(client)

	if err := program.Run(*category); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
