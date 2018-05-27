package db

import "github.com/whiter4bbit/imdb-datasets/datasets"

//go:generate msgp
type DBEntry struct {
	Movie   *datasets.MovieIdentity
	Episode *datasets.EpisodeIdentity
	SeqId   uint32
}
