package movies

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ChristianLapinig/tmdb-cli/constants"
)

func FetchMovies(category, accessToken string) {
	url := fmt.Sprintf("%s/authentication", constants.BaseURL)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
}
