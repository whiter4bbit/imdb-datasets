package datasets

import (
	"fmt"
	"github.com/tinylib/msgp/msgp"
)

type AttrType byte

const (
	AttrActors AttrType = iota
	AttrGenres
	AttrPlots
	AttrRatings
)

type Attribute interface {
	msgp.Marshaler
	Update(*Attributes)
	GetType() AttrType
}

//go:generate msgp
type Attributes struct {
	Plot   string
	Rating float64
	Votes  int
	Actors []string
	Genres []string
}

func UnmarshalAttribute(t AttrType, b []byte) (Attribute, error) {
	switch t {
	case AttrActors:
		actors := &Actors{}
		_, err := actors.UnmarshalMsg(b)
		return actors, err
	case AttrGenres:
		genres := &Genres{}
		_, err := genres.UnmarshalMsg(b)
		return genres, err
	case AttrPlots:
		plots := &Plots{}
		_, err := plots.UnmarshalMsg(b)
		return plots, err
	case AttrRatings:
		ratings := &Ratings{}
		_, err := ratings.UnmarshalMsg(b)
		return ratings, err
	}
	panic(fmt.Sprintf("unknown type %d", t))
}

type Actors struct {
	Actor string
}

func (a *Actors) Update(attrs *Attributes) {
	attrs.Actors = append(attrs.Actors, a.Actor)
}

func (a *Actors) GetType() AttrType {
	return AttrActors
}

type Genres struct {
	Genre string
}

func (g *Genres) Update(m *Attributes) {
	m.Genres = append(m.Genres, g.Genre)
}

func (g *Genres) GetType() AttrType {
	return AttrGenres
}

type Plots struct {
	Plot string
}

func (p *Plots) Update(m *Attributes) {
	m.Plot = p.Plot
}

func (p *Plots) GetType() AttrType {
	return AttrPlots
}

type Ratings struct {
	Votes  int
	Rating float64
}

func (r *Ratings) Update(m *Attributes) {
	m.Votes = r.Votes
	m.Rating = r.Rating
}

func (r *Ratings) GetType() AttrType {
	return AttrRatings
}
