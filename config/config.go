package config

import "os"

type Config struct {
	AccessToken string
}

func InitializeConfig() (*Config, error) {
	accessToken := os.Getenv("TMDB_API_READ_ACCESS_TOKEN")
	config := &Config{
		AccessToken: accessToken,
	}
	return config, nil
}
