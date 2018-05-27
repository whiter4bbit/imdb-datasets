package main

import (
	"flag"
	"fmt"
	"github.com/whiter4bbit/imdb-datasets/datasets/legacy"
	"github.com/whiter4bbit/imdb-datasets/db"
	"os"
)

func exportMain(args []string) {
	flagSet := flag.NewFlagSet("export", flag.ExitOnError)

	flagSet.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flagSet.PrintDefaults()
	}

	dbPath := flagSet.String("db-path", "imdb.db", "output path")

	fstPath := flagSet.String("fst-path", "titles.fst", "fst output path")

	imdbPath := flagSet.String("imdb-path", "ftp://ftp.fu-berlin.de/pub/misc/movies/database/frozendata", "imdb path")

	flagSet.Parse(args)

	if err := legacy.Export(*imdbPath, *dbPath, *fstPath); err != nil {
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

	fstPath := flagSet.String("fst-path", "titles.fst", "path to fst")

	dist := flagSet.Int("dist", 2, "levenshtein distance")

	query := flagSet.String("query", "rain man", "query")

	flagSet.Parse(args)

	reader, err := db.NewReader(*dbPath)
	if err != nil {
		fmt.Printf("Can not create reader: %v\n", err)
		os.Exit(2)
	}
	defer reader.Close()

	search, err := db.NewSearch(*fstPath, reader, *dist)
	if err != nil {
		fmt.Printf("Can not create search: %v\n", err)
		os.Exit(2)
	}

	results, err := search.Search(*query)
	if err != nil {
		fmt.Printf("Can search: %v\n", err)
		os.Exit(2)
	}

	for _, result := range results {
		fmt.Printf("%s\n", result.ShortString())
	}
}

func printUsageAndExit() {
	fmt.Printf(`Usage: %s [command]
Command:
  - export
  - search
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
	default:
		printUsageAndExit()
	}
}
