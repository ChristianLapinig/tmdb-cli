package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://api.themoviedb.org/3"

func main() {
	accessToken := os.Getenv("TMDB_API_READ_ACCESS_TOKEN")
	authHeaderVal := fmt.Sprintf("Bearer %s", accessToken)
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/authentication", baseURL), nil)
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
