package datasets

import "fmt"

type Movie struct {
	Id Hash
	*MovieIdentity
	*Attributes
	Episodes []*Episode
}

func (m *Movie) ShortString() string {
	return fmt.Sprintf("{Title: %q, Year: %d, Actors: %+v, Genres: %+v, Plot: %s}", m.Title, m.MovieIdentity.Year, m.Actors, m.Genres, m.Plot)
}

type Episode struct {
	Id Hash
	*EpisodeIdentity
	*Attributes
}
