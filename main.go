package main

import (
	"os"

	"github.com/ChristianLapinig/tmdb-cli/program"
)

func main() {
	accessToken := os.Getenv("TMDB_API_READ_ACCESS_TOKEN")
	program.Program(accessToken)
}
