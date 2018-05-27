package db

import (
	"github.com/couchbase/vellum"
	"github.com/couchbase/vellum/levenshtein"
	"github.com/whiter4bbit/imdb-datasets/datasets"
)

type Search struct {
	fst    *vellum.FST
	reader *Reader
	dist   int
}

func NewSearch(fstPath string, reader *Reader, dist int) (*Search, error) {
	fst, err := vellum.Open(fstPath)
	if err != nil {
		return nil, err
	}

	return &Search{
		fst:    fst,
		reader: reader,
		dist:   dist,
	}, nil
}

func (s *Search) Search(query string) ([]*datasets.Movie, error) {
	lst, err := levenshtein.New(query, s.dist)
	if err != nil {
		return nil, err
	}

	var seqIds []uint32

	iter, err := s.fst.Search(lst, nil, nil)
	for err == nil {
		_, id := iter.Current()

		seqIds = append(seqIds, uint32(id))

		err = iter.Next()
	}

	return s.reader.GetByIds(seqIds)
}
