package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/whiter4bbit/imdb-datasets/datasets"
	"github.com/whiter4bbit/imdb-datasets/datasets/legacy"
	"github.com/whiter4bbit/imdb-datasets/db"
	"math/rand"
	"os"
)

func exportMain(args []string) {
	flagSet := flag.NewFlagSet("export", flag.ExitOnError)

	flagSet.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flagSet.PrintDefaults()
	}

	dbPath := flagSet.String("db-path", "imdb.db", "output path")

	maxWritesPerTx := flagSet.Int("max-writes-per-tx", 50000, "max writes per transaction")

	imdbPath := flagSet.String("imdb-path", "ftp://ftp.fu-berlin.de/pub/misc/movies/database/frozendata", "imdb path")

	flagSet.Parse(args)

	if err := legacy.Export(*imdbPath, *dbPath, *maxWritesPerTx); err != nil {
		fmt.Printf("Can not export datasets: %v\n", err)
		os.Exit(2)
	}
}

func searchMain(args []string) {
	flagSet := flag.NewFlagSet("search", flag.ExitOnError)

	flagSet.Usage = func() {
		fmt.Printf("Usage of %s\n", os.Args[0])
		flagSet.PrintDefaults()
	}

	dbPath := flagSet.String("db-path", "imdb.db", "path to movie database")

	query := flagSet.String("query", "rain man", "query")

	flagSet.Parse(args)

	reader, err := db.NewReader(*dbPath)
	if err != nil {
		fmt.Printf("Can not create reader: %v\n", err)
		os.Exit(2)
	}
	defer reader.Close()

	results, err := reader.GetByTitlePrefix([]byte(*query), nil)
	if err != nil {
		fmt.Printf("Can search: %v\n", err)
		os.Exit(2)
	}

	for _, result := range results {
		fmt.Printf("%s\n", result.ShortString())
	}
}

func dumpMain(args []string) {
	flagSet := flag.NewFlagSet("dump", flag.ExitOnError)

	flagSet.Usage = func() {
		fmt.Printf("Usage of %s\n", os.Args[0])
		flagSet.PrintDefaults()
	}

	dbPath := flagSet.String("db-path", "imdb.db", "path to movie database")

	sample := flagSet.Float64("sample-rate", 1.0, "same rate")

	rating := flagSet.Float64("rating", 8.0, "rating lower bound")

	flagSet.Parse(args)

	reader, err := db.NewReader(*dbPath)
	if err != nil {
		fmt.Printf("Can not create reader: %v\n", err)
		os.Exit(2)
	}
	defer reader.Close()

	if err := reader.ForEachMovie(func(movie *datasets.Movie) error {
		if len(movie.Plot) == 0 {
			return nil
		}

		if movie.Rating < *rating {
			return nil
		}

		if rand.Float64() > *sample {
			return nil
		}

		movie.Episodes = nil

		b, err := json.Marshal(movie)
		if err != nil {
			return err
		}

		fmt.Printf(string(b) + "\n")

		return nil
	}); err != nil {
		fmt.Printf("Can not dump as json: %v\n", err)
		os.Exit(2)
	}
}

func printUsageAndExit() {
	fmt.Printf(`Usage: %s [command]
Command:
  - export
  - search
  - dump
`, os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) == 1 {
		printUsageAndExit()
	}

	switch os.Args[1] {
	case "export":
		exportMain(os.Args[2:])
	case "search":
		searchMain(os.Args[2:])
	case "dump":
		dumpMain(os.Args[2:])
	default:
		printUsageAndExit()
	}
}
