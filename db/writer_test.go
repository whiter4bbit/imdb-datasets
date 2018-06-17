package db

import (
	"github.com/stretchr/testify/require"
	"github.com/whiter4bbit/imdb-datasets/datasets"
	"io/ioutil"
	"os"
	"testing"
)

func TestWriter(t *testing.T) {
	path, err := ioutil.TempFile(".", "db")
	require.Nil(t, err)

	defer os.RemoveAll(path.Name())

	writer, err := NewWriter(path.Name(), 10000)
	require.Nil(t, err)

	data := []struct {
		movie   *datasets.MovieIdentity
		episode *datasets.EpisodeIdentity
		attrs   []datasets.Attribute
	}{
		{
			movie: &datasets.MovieIdentity{
				Title: "Knockin' on Heaven's Door",
				Year:  1997,
			},
			attrs: []datasets.Attribute{
				&datasets.Actors{Actor: "Till Scheiger"},
				&datasets.Actors{Actor: "Jan Josef Liefers"},
				&datasets.Genres{Genre: "drama"},
				&datasets.Genres{Genre: "comedy"},
			},
		},
		{
			movie: &datasets.MovieIdentity{
				Title: "Rain Man",
				Year:  1988,
			},
			attrs: []datasets.Attribute{
				&datasets.Actors{Actor: "Dustin Hoffman"},
				&datasets.Actors{Actor: "Tom Cruise"},
			},
		},
		{
			movie: &datasets.MovieIdentity{
				Title: "Simpsons",
				Year:  1999,
			},
			attrs: []datasets.Attribute{
				&datasets.Actors{Actor: "Homer"},
				&datasets.Actors{Actor: "Marge"},
			},
		},
		{
			episode: &datasets.EpisodeIdentity{
				Title:         "Simpsons",
				Year:          1999,
				EpisodeTitle:  "Episode with Bart",
				EpisodeNumber: 1,
				SeassonNumber: 1,
			},
			attrs: []datasets.Attribute{
				&datasets.Actors{Actor: "Bart"},
			},
		},
		{
			episode: &datasets.EpisodeIdentity{
				Title:         "Simpsons",
				Year:          1999,
				EpisodeTitle:  "Episode with Liza",
				EpisodeNumber: 1,
				SeassonNumber: 2,
			},
			attrs: []datasets.Attribute{
				&datasets.Actors{Actor: "Liza"},
			},
		},
		{
			episode: &datasets.EpisodeIdentity{
				Title:         "Simpsons",
				Year:          1999,
				EpisodeTitle:  "Episode with Milhouse",
				EpisodeNumber: 1,
				SeassonNumber: 3,
			},
			attrs: []datasets.Attribute{
				&datasets.Actors{Actor: "Milhouse"},
			},
		},
	}

	for _, entry := range data {
		var hash datasets.Hash

		if entry.movie != nil {
			require.Nil(t, writer.WriteMovie(entry.movie))
			hash = entry.movie.MakeHash()
		} else {
			require.Nil(t, writer.WriteEpisode(entry.episode))
			hash = entry.episode.MakeHash()
		}

		for _, attr := range entry.attrs {
			require.Nil(t, writer.WriteAttribute(hash, attr))
		}
	}

	require.Nil(t, writer.Close())

	reader, err := NewReader(path.Name())
	require.Nil(t, err)

	GetMovie := func(n int) *datasets.Movie {
		var hash datasets.Hash
		if data[n].movie != nil {
			hash = data[n].movie.MakeHash()
		} else {
			hash = data[n].episode.MakeParentHash()
		}

		movies, err := reader.GetByHash([]datasets.Hash{hash})
		require.Nil(t, err)
		require.Equal(t, 1, len(movies))
		require.Equal(t, hash, movies[0].Id)

		var empty datasets.Hash

		movies[0].Id = empty

		for _, episode := range movies[0].Episodes {
			episode.Id = empty
		}

		return movies[0]
	}

	require.Equal(t, GetMovie(0), &datasets.Movie{
		MovieIdentity: &datasets.MovieIdentity{
			Title: "Knockin' on Heaven's Door",
			Year:  1997,
		},
		Attributes: &datasets.Attributes{
			Actors: []string{
				"Till Scheiger",
				"Jan Josef Liefers",
			},
			Genres: []string{
				"drama",
				"comedy",
			},
		},
	})

	require.Equal(t, GetMovie(2), &datasets.Movie{
		MovieIdentity: &datasets.MovieIdentity{
			Title: "Simpsons",
			Year:  1999,
		},
		Attributes: &datasets.Attributes{
			Actors: []string{
				"Homer",
				"Marge",
			},
		},
		Episodes: []*datasets.Episode{
			&datasets.Episode{
				EpisodeIdentity: &datasets.EpisodeIdentity{
					Title:         "Simpsons",
					Year:          1999,
					EpisodeTitle:  "Episode with Liza",
					EpisodeNumber: 1,
					SeassonNumber: 2,
				},
				Attributes: &datasets.Attributes{
					Actors: []string{
						"Liza",
					},
				},
			},
			&datasets.Episode{
				EpisodeIdentity: &datasets.EpisodeIdentity{
					Title:         "Simpsons",
					Year:          1999,
					EpisodeTitle:  "Episode with Bart",
					EpisodeNumber: 1,
					SeassonNumber: 1,
				},
				Attributes: &datasets.Attributes{
					Actors: []string{
						"Bart",
					},
				},
			},
			&datasets.Episode{
				EpisodeIdentity: &datasets.EpisodeIdentity{
					Title:         "Simpsons",
					Year:          1999,
					EpisodeTitle:  "Episode with Milhouse",
					EpisodeNumber: 1,
					SeassonNumber: 3,
				},
				Attributes: &datasets.Attributes{
					Actors: []string{
						"Milhouse",
					},
				},
			},
		},
	})
}
