package legacy

import (
    "github.com/stretchr/testify/require"
    "github.com/whiter4bbit/imdb-datasets/datasets"
    "github.com/whiter4bbit/imdb-datasets/db"
    "os"
    "testing"
)

func TestExportActors(t *testing.T) {
    reader, err := os.Open("testdata/actors.list")
    require.Nil(t, err)
    defer reader.Close()

    writer := &db.MockWriter{}

    require.Nil(t, exportActors(reader, writer))

    require.Equal(t, []*db.HashAndAttribute{
        {(&datasets.MovieIdentity{Title: "Fatherhood 101", Year: 2013}).MakeHash(), &datasets.Actors{Actor: "$lim, Bee Moe"}},
        {(&datasets.MovieIdentity{Title: "For Thy Love 2", Year: 2009}).MakeHash(), &datasets.Actors{Actor: "$lim, Bee Moe"}},
        {(&datasets.MovieIdentity{Title: "Tria33", Year: 2015}).MakeHash(), &datasets.Actors{Actor: "'77"}},
        {(&datasets.EpisodeIdentity{Title: "Simpsons", Year: 1999, SeassonNumber: 1, EpisodeNumber: 2}).MakeHash(), &datasets.Actors{Actor: "'77"}},
    }, writer.Attrs)
}

func TestExportGenres(t *testing.T) {
    reader, err := os.Open("testdata/genres.list")
    require.Nil(t, err)
    defer reader.Close()

    writer := &db.MockWriter{}

    require.Nil(t, exportGenres(reader, writer))

    require.Equal(t, []*db.HashAndAttribute{
        {(&datasets.MovieIdentity{Title: "!Next?", Year: 1994}).MakeHash(), &datasets.Genres{Genre: "Documentary"}},
        {(&datasets.MovieIdentity{Title: "#1 Single", Year: 2006}).MakeHash(), &datasets.Genres{Genre: "Reality-TV"}},
        {(&datasets.MovieIdentity{Title: "#15SecondScare", Year: 2015}).MakeHash(), &datasets.Genres{Genre: "Horror"}},
    }, writer.Attrs)
}

func TestExportRatings(t *testing.T) {
    reader, err := os.Open("testdata/ratings.list")
    require.Nil(t, err)
    defer reader.Close()

    writer := &db.MockWriter{}

    require.Nil(t, exportRatings(reader, writer))

    require.Equal(t, []*db.HashAndAttribute{
        {(&datasets.MovieIdentity{Title: "The Shawshank Redemption", Year: 1994}).MakeHash(), &datasets.Ratings{Rating: 9.2, Votes: 1786238}},
        {(&datasets.MovieIdentity{Title: "The Godfather", Year: 1972}).MakeHash(), &datasets.Ratings{Rating: 9.2, Votes: 1219308}},
        {(&datasets.MovieIdentity{Title: "The Godfather: Part II", Year: 1974}).MakeHash(), &datasets.Ratings{Rating: 9.0, Votes: 838449}},
    }, writer.Attrs)
}

func TestExportPlots(t *testing.T) {
    reader, err := os.Open("testdata/plots.list")
    require.Nil(t, err)
    defer reader.Close()

    writer := &db.MockWriter{}

    require.Nil(t, exportPlots(reader, writer))

    require.Equal(t, []*db.HashAndAttribute{
        {(&datasets.MovieIdentity{Title: "#7DaysLater", Year: 2013}).MakeHash(), &datasets.Plots{Plot: `#7dayslater is an interactive comedy series featuring an ensemble cast of YouTube celebrities. Each week the audience writes the brief via social media for an all-new episode featuring a well-known guest-star. Seven days later that week's episode premieres on TV and across multiple platforms.`}},
        {(&datasets.MovieIdentity{Title: "#BlackLove", Year: 2015}).MakeHash(), &datasets.Plots{Plot: `This week, the five women work on getting what they want in a relationship by being more open and learning to communicate properly. Cynthia, after herself to Karl.`}},
    }, writer.Attrs)
}

func TestExportMovies(t *testing.T) {
    reader, err := os.Open("testdata/movies.list")
    require.Nil(t, err)
    defer reader.Close()

    writer := &db.MockWriter{}

    require.Nil(t, exportMovies(reader, writer))

    require.Equal(t, []*datasets.MovieIdentity{
        &datasets.MovieIdentity{Title: "!Next?", Year: 1994},
        &datasets.MovieIdentity{Title: "#1 Single", Year: 2006},
        &datasets.MovieIdentity{Title: "#15SecondScare", Year: 2015},
    }, writer.Movies)

    require.Equal(t, []*datasets.EpisodeIdentity{
        &datasets.EpisodeIdentity{Title: "#1 Single", Year: 2006, EpisodeTitle: "Cats and Dogs", EpisodeNumber: 4, SeassonNumber: 1},
        &datasets.EpisodeIdentity{Title: "#1 Single", Year: 2006, EpisodeTitle: "Finishing a Chapter", EpisodeNumber: 5, SeassonNumber: 1},
    }, writer.Episodes)
}

func TestMovieIdParse(t *testing.T) {
    inputs := []struct {
        expr     string
        expected *parsedIdentity
    }{
        {
            expr: `"#15SecondScare" (2015)`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "#15SecondScare",
                    Year:  2015,
                },
            },
        },
        {
            expr: `"#15SecondScare" (2015) {Because We Don't Want You to Fall Asleep (#1.3)}`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "#15SecondScare",
                    Year:  2015,
                },
                episode: &datasets.EpisodeIdentity{
                    Title:         "#15SecondScare",
                    Year:          2015,
                    EpisodeTitle:  "Because We Don't Want You to Fall Asleep",
                    EpisodeNumber: 3,
                    SeassonNumber: 1,
                },
            },
        },
        {
            expr: `"#15SecondScare" (2015) {Beauty Wrap}`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "#15SecondScare",
                    Year:  2015,
                },
                episode: &datasets.EpisodeIdentity{
                    Title:        "#15SecondScare",
                    Year:         2015,
                    EpisodeTitle: "Beauty Wrap",
                },
            },
        },
        {
            expr: `"(Des)encontros" (2014)  2014-????`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "(Des)encontros",
                    Year:  2014,
                },
            },
        },
        {
            expr: `"#CandidlyNicole" (2013) {Day with Dad (Special Guest: Lionel Richie) (#1.7)}  ????`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "#CandidlyNicole",
                    Year:  2013,
                },
                episode: &datasets.EpisodeIdentity{
                    Title:         "#CandidlyNicole",
                    Year:          2013,
                    EpisodeTitle:  "Day with Dad (Special Guest: Lionel Richie)",
                    SeassonNumber: 1,
                    EpisodeNumber: 7,
                },
            },
        },
        {
            expr: `"#Nofilter" (2014) {Pilot (#1.1)} {{SUSPENDED}}         2014`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "#Nofilter",
                    Year:  2014,
                },
                episode: &datasets.EpisodeIdentity{
                    Title:         "#Nofilter",
                    Year:          2014,
                    EpisodeTitle:  "Pilot",
                    SeassonNumber: 1,
                    EpisodeNumber: 1,
                },
            },
        },
        {
            expr: `"#Cocinando" (2015) {{SUSPENDED}}                       2015-????`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "#Cocinando",
                    Year:  2015,
                },
            },
        },
        {
            expr: `"Jura" (2006) {(#1.30)}  [Francisco]  <14>`,
            expected: &parsedIdentity{
                movie: &datasets.MovieIdentity{
                    Title: "Jura",
                    Year:  2006,
                },
                episode: &datasets.EpisodeIdentity{
                    Title:         "Jura",
                    Year:          2006,
                    SeassonNumber: 1,
                    EpisodeNumber: 30,
                },
            },
        },
    }

    for _, input := range inputs {
        parsed, err := parseIdentity(input.expr)
        if err != nil {
            require.Nil(t, err)
        }

        require.Equal(t, input.expected, parsed)
    }
}
