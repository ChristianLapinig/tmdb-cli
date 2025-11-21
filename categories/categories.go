package categories

type Category string

const (
	NowPlaying Category = "now_playing"
	Popular    Category = "popular"
	TopRated   Category = "top_rated"
	Upcoming   Category = "upcoming"
)

var Categories = []Category{
	NowPlaying,
	Popular,
	Upcoming,
	TopRated,
}
