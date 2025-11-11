package main

import (
	"fmt"
	"os"

	"github.com/ChristianLapinig/tmdb-cli/config"
	"github.com/ChristianLapinig/tmdb-cli/program"
)

func main() {
	config, err := config.InitializeConfig()
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
	program.Program(config.AccessToken)
}
