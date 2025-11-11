package program

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ChristianLapinig/tmdb-cli/constants"
)

func Program(accessToken string) {
	if accessToken == "" {
		fmt.Println(constants.AccessTokenRequired)
		os.Exit(1)
	}

	authHeaderVal := fmt.Sprintf("Bearer %s", accessToken)
	url := fmt.Sprintf("%s/authentication", constants.BaseURL)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authHeaderVal)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
