package main

import (
	"fmt"

	"github.com/ChristianLapinig/tmdb-cli/config"
	"github.com/ChristianLapinig/tmdb-cli/program"
)

func main() {
	config, err := config.InitializeConfig()
	if err != nil {
		fmt.Println("ERROR:", err.Error())
	}
	program.Program(config.AccessToken)
}
