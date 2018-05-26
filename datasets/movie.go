package datasets

type MovieExtent interface {
	MovieId() string
}

type Movie struct {
	MovieId string
	Title   string
	Plot    string
	Rating  float64
	Year    int
	Actors  []string
}
