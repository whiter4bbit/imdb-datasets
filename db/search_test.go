package db

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"github.com/whiter4bbit/imdb-datasets/datasets"
	"io/ioutil"
	"os"
	"sort"
	"testing"
)

func TestSearch(t *testing.T) {
	dbPath, err := ioutil.TempFile(".", "db")
	require.Nil(t, err)

	fstPath, err := ioutil.TempFile(".", "fst")
	require.Nil(t, err)

	defer os.RemoveAll(dbPath.Name())
	defer os.RemoveAll(fstPath.Name())

	writer, err := NewWriter(dbPath.Name(), 10000)
	require.Nil(t, err)

	movies := []*datasets.MovieIdentity{
		{
			Title: "Friends",
			Year:  1997,
		},
		{
			Title: "Friends",
			Year:  2001,
		},
		{
			Title: "Simpsons",
			Year:  1999,
		},
		{
			Title: "Simpsons",
			Year:  1989,
		},
	}

	for _, movie := range movies {
		require.Nil(t, writer.WriteMovie(movie))
	}

	require.Nil(t, writer.Close())

	reader, err := NewReader(dbPath.Name())
	require.Nil(t, err)

	fstWriter, err := NewFSTWriter(fstPath.Name())
	require.Nil(t, err)
	require.Nil(t, reader.ForEachTitleIndex(fstWriter.Write))
	fstWriter.Close()

	search, err := NewSearch(fstPath.Name(), reader, 0)
	require.Nil(t, err)

	Sort := func(hashes []datasets.Hash) []datasets.Hash {
		sort.Slice(hashes, func(i, j int) bool {
			return bytes.Compare(hashes[i][:], hashes[j][:]) < 0
		})
		return hashes
	}

	SearchHashes := func(query string) []datasets.Hash {
		movies, err := search.Search(query)
		require.Nil(t, err)

		var hashes []datasets.Hash
		for _, movie := range movies {
			hashes = append(hashes, movie.MovieIdentity.MakeHash())
		}

		return Sort(hashes)
	}

	require.Equal(t, Sort([]datasets.Hash{movies[0].MakeHash(), movies[1].MakeHash()}), SearchHashes("friends"))
	require.Equal(t, Sort([]datasets.Hash{movies[2].MakeHash(), movies[3].MakeHash()}), SearchHashes("simpsons"))
}
