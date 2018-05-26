package main

import (
	"flag"
	"fmt"
	"github.com/whiter4bbit/imdb-datasets/datasets"
	"os"
	"strings"
)

func exportMain(args []string) {
	flagSet := flag.NewFlagSet("join", flag.ExitOnError)

	flagSet.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flagSet.PrintDefaults()
	}

	datasetsToJoin := flagSet.String("datasets", "actors", "list of datasets to export")

	outputPath := flagSet.String("out", ".", "output path")

	imdbPath := flagSet.String("imdb-path", "", "imdb path")

	flagSet.Parse(args)

	if err := datasets.Export(*imdbPath, *outputPath, strings.Split(*datasetsToJoin, ",")); err != nil {
		fmt.Printf("Can not export datasets: %v\n", err)
		os.Exit(2)
	}
}

func printUsageAndExit() {
	fmt.Printf(`Usage: %s [command]
Command:
  - export
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
	default:
		printUsageAndExit()
	}
}
