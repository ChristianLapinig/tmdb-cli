package models

type Movie struct {
	ID            int     `json:"id"`
	OriginalTitle string  `json:"original_title"`
	Overview      string  `json:"overview"`
	Language      string  `json:"original_language"`
	ReleaseDate   string  `json:"release_date"`
	Rating        float64 `json:"vote_average"`
}
