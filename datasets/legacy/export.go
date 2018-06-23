package legacy

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/whiter4bbit/imdb-datasets/datasets"
	"github.com/whiter4bbit/imdb-datasets/db"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	knownDatasets = []struct {
		name string
		f    func(io.ReadCloser, db.Writer) error
	}{
		{"movies", exportMovies},
		{"actors", exportActors},
		{"actresses", exportActors},
		{"genres", exportGenres},
		{"plot", exportPlots},
		{"ratings", exportRatings},
	}

	ftpPref = "ftp://"
)

func canNotParseLine(line string) error {
	return fmt.Errorf("can not parse line %q", line)
}

func Export(imdbPath, dbPath string, maxWritesPerTx int) error {
	writer, err := db.NewWriter(dbPath, maxWritesPerTx)
	if err != nil {
		return err
	}

	var conn *ftp.ServerConn

	if strings.HasPrefix(imdbPath, ftpPref) {
		segments := strings.SplitN(imdbPath[len(ftpPref):], "/", 2)

		address := segments[0]

		if strings.Index(address, ":") == -1 {
			address += ":21"
		}

		conn, err = ftp.Connect(address)
		if err != nil {
			return err
		}

		if err := conn.Login("guest", "guest"); err != nil {
			return err
		}

		if len(segments) == 2 {
			if err := conn.ChangeDir(segments[1]); err != nil {
				return err
			}
		}
	}

	log.Printf("Exporting db to: %s", dbPath)

	for _, dataset := range knownDatasets {
		var f io.ReadCloser

		listPath := dataset.name + ".list.gz"

		if conn != nil {
			f, err = conn.Retr(listPath)
		} else {
			f, err = os.Open(imdbPath + string(os.PathSeparator) + listPath)
		}

		if err != nil {
			return err
		}

		r, err := gzip.NewReader(f)
		if err != nil {
			return err
		}

		if err := dataset.f(r, writer); err != nil {
			return err
		}

		f.Close()

		log.Printf("Dataset exported: %s", dataset.name)
	}

	if conn != nil {
		conn.Quit()
	}

	return writer.Close()
}

func exportLines(r io.ReadCloser, f func(line string) error) error {
	var (
		moviesNotFound int
		headerSeen     bool
	)

	decoder := charmap.ISO8859_1.NewDecoder().Reader(r)

	scanner := bufio.NewScanner(decoder)

	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "TOP 250 MOVIES (25000+ VOTES)") {
			break
		}

		if !headerSeen && strings.HasSuffix(scanner.Text(), " LIST") {
			headerSeen = true
			continue
		}

		if headerSeen && strings.HasPrefix(scanner.Text(), "=========") {
			break
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		err := f(line)
		if err == db.NotFoundErr {
			moviesNotFound += 1
			continue
		} else if err != nil {
			return err
		}

	}

	log.Printf("Movies not found: %d", moviesNotFound)

	return scanner.Err()
}

func exportActors(r io.ReadCloser, writer db.Writer) error {
	var (
		eof          bool
		currentActor *datasets.Actors
	)

	return exportLines(r, func(line string) error {
		if eof || strings.HasPrefix(line, "------------") {
			eof = true
			return nil
		}

		attrs := strings.SplitN(line, "\t", 2)

		if len(attrs) == 2 {
			actor := strings.Trim(attrs[0], "\t ")
			title := strings.Trim(attrs[1], "\t ")

			if actor == "Name" || actor == "----" {
				return nil
			}

			if hash, err := readHash(title); err != nil {
				return err
			} else if len(actor) == 0 {
				return writer.WriteAttribute(hash, currentActor)
			} else {
				currentActor = &datasets.Actors{Actor: actor}
				return writer.WriteAttribute(hash, currentActor)
			}
		}

		return nil
	})
}

func exportGenres(r io.ReadCloser, writer db.Writer) error {
	return exportLines(r, func(line string) error {
		attrs := strings.SplitN(line, "\t", 2)
		if len(attrs) == 2 {
			if hash, err := readHash(strings.Trim(attrs[0], "\t ")); err != nil {
				return err
			} else {
				genre := strings.Trim(attrs[1], "\t ")
				return writer.WriteAttribute(hash, &datasets.Genres{Genre: genre})
			}
		}

		return canNotParseLine(line)
	})
}

func exportRatings(r io.ReadCloser, writer db.Writer) error {
	var (
		eof       bool
		startData bool
	)

	return exportLines(r, func(line string) error {
		if eof || strings.HasPrefix(line, "------------") {
			eof = true
			return nil
		}

		if strings.HasPrefix(line, "BOTTOM 10 MOVIES (1500+ VOTES)") ||
			strings.HasPrefix(line, "MOVIE RATINGS REPORT") ||
			strings.HasPrefix(line, "------------") {

			startData = false
			return nil
		}

		if strings.HasPrefix(line, "New  Distribution  Votes  Rank  Title") {
			startData = true
			return nil
		}

		if !startData {
			return nil
		}

		fields := strings.Split(line, " ")

		var parsedFields []string

		for len(fields) > 0 && len(parsedFields) < 3 {
			if len(fields[0]) > 0 {
				parsedFields = append(parsedFields, fields[0])
			}

			fields = fields[1:]
		}

		if len(parsedFields) != 3 {
			return canNotParseLine(line)
		}

		title := strings.Trim(strings.Join(fields, " "), " ")

		votes, err := strconv.Atoi(parsedFields[1])
		if err != nil {
			return canNotParseLine(line)
		}

		rating, err := strconv.ParseFloat(parsedFields[2], 64)
		if err != nil {
			return canNotParseLine(line)
		}

		if hash, err := readHash(title); err != nil {
			return err
		} else {
			return writer.WriteAttribute(hash, &datasets.Ratings{
				Rating: rating,
				Votes:  votes,
			})
		}
	})
}

var (
	mv = "MV:"
	pl = "PL:"
)

func exportPlots(r io.ReadCloser, writer db.Writer) error {
	var (
		hash datasets.Hash
		plot string
	)

	flush := func() error {
		if plot != "" {
			err := writer.WriteAttribute(hash, &datasets.Plots{Plot: strings.Trim(plot, " ")})
			if err == db.NotFoundErr {
				plot = ""
				return nil
			} else if err != nil {
				plot = ""
				return err
			}

			plot = ""
		}
		return nil
	}

	if err := exportLines(r, func(line string) error {
		if strings.HasPrefix(line, mv) {
			if err := flush(); err != nil {
				return err
			}

			if h, err := readHash(line[len(mv):]); err != nil {
				return err
			} else {
				hash = h
			}
		}

		if strings.HasPrefix(line, pl) {
			plot += line[len(pl):]
		}

		return nil
	}); err != nil {
		return err
	}

	err := flush()

	if err == db.NotFoundErr {
		return nil
	}

	return err
}

func exportMovies(r io.ReadCloser, writer db.Writer) error {
	var eof bool
	return exportLines(r, func(line string) error {
		if eof || strings.HasPrefix(line, "------------") {
			eof = true
			return nil
		}

		if identity, err := parseIdentity(line); err != nil {
			return err
		} else {
			if identity.episode != nil {
				if err := writer.WriteEpisode(identity.episode); err == nil {
					return nil
				} else if err != db.NotFoundErr {
					log.Printf("can not write episode: %+v: %+v", identity.episode, err)
					return err
				}
			}

			if identity.movie != nil {
				if err := writer.WriteMovie(identity.movie); err != nil {
					log.Printf("can not write movie: %+v: %+v", identity.movie, err)
					return err
				}
			}

			if identity.episode != nil {
				if err := writer.WriteEpisode(identity.episode); err != nil {
					log.Printf("can not write episode: %+v: %+v", identity.episode, err)
					return nil
				}
			}

		}
		return nil
	})
}

var (
	movieTitleRE = regexp.MustCompile(`^(.*?)(\([\d\?]{4}\/?[XLVI]*\))( *(?:\(TV\)|\(V\)|\(VG\))? *(\{[^\{}]*\})?).*`)
	episodeRE    = regexp.MustCompile(`^\{(.*?)(\(\#\d+\.\d+\))?\}$`)
)

var spaceAndQuotes = `" `

type parsedIdentity struct {
	movie   *datasets.MovieIdentity
	episode *datasets.EpisodeIdentity
}

// "#15SecondScare" (2015) {Because We Don't Want You to Fall Asleep (#1.3)}
func parseIdentity(expr string) (*parsedIdentity, error) {

	noMatch := func(cause string) error {
		return fmt.Errorf(cause+": `%s`", expr)
	}

	submatch := movieTitleRE.FindStringSubmatch(expr)

	if len(submatch) == 0 {
		return nil, noMatch("does not match the pattern")
	}

	movie := &datasets.MovieIdentity{}

	movie.Title = strings.Trim(submatch[1], spaceAndQuotes)

	if year, err := strconv.Atoi(strings.Trim(submatch[2], `(VLVI/?)`)); err == nil {
		movie.Year = year
	}

	if len(submatch[4]) > 0 {
		episodeSubmatch := episodeRE.FindStringSubmatch(submatch[4])

		if len(episodeSubmatch) == 0 {
			return nil, noMatch("can not match episode segment")
		}

		var (
			seassonNumber int
			episodeNumber int
			err           error
		)

		episodeTitle := strings.Trim(episodeSubmatch[1], spaceAndQuotes)

		episodeNoExpr := strings.Trim(episodeSubmatch[2], `(#)`)

		if seassonAndNo := strings.Split(episodeNoExpr, "."); len(seassonAndNo) == 2 {
			if seassonNumber, err = strconv.Atoi(seassonAndNo[0]); err != nil {
				return nil, noMatch("can not parse seasson no: `" + seassonAndNo[0] + "`")
			}

			if episodeNumber, err = strconv.Atoi(seassonAndNo[1]); err != nil {
				return nil, noMatch("can not parse episode no")
			}
		}

		return &parsedIdentity{
			movie: movie,
			episode: &datasets.EpisodeIdentity{
				Title:         movie.Title,
				Year:          movie.Year,
				EpisodeTitle:  episodeTitle,
				EpisodeNumber: episodeNumber,
				SeassonNumber: seassonNumber,
			},
		}, nil
	}

	return &parsedIdentity{
		movie: movie,
	}, nil
}

func readHash(expr string) (datasets.Hash, error) {
	id, err := parseIdentity(expr)
	if err != nil {
		var empty [datasets.HashSize]byte
		return empty, err
	}

	if id.episode != nil {
		return id.episode.MakeHash(), nil
	}

	return id.movie.MakeHash(), nil
}
