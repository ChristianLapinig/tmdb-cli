package program

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ChristianLapinig/tmdb-cli/categories"
	"github.com/ChristianLapinig/tmdb-cli/constants"
	//"github.com/ChristianLapinig/tmdb-cli/movies"
)

func Program(accessToken string) {
	flag.Usage = func() {
		fmt.Println("Usage: tmdb-cli --type <category>")
		fmt.Println("Please note that the --type flag must also be passed.")
		fmt.Printf("One of the following categories must be passed as an argument:\n")
		for _, category := range categories.Categories {
			fmt.Println(strings.Repeat(" ", 4) + "- " + string(category))
		}
	}

	if accessToken == "" {
		fmt.Println(constants.AccessTokenRequired)
		os.Exit(1)
	}

	typeFlag := flag.Bool("type", false, "The category of movies to be display (i.e, now_playing, popular, upcoming, and top_rated")

	flag.Parse()

	if *typeFlag {
		category := flag.Args()[0]
		if category == "" {
			fmt.Println("No category specified. Please see usage below")
		}
	}
}
