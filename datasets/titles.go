package datasets

import (
	"crypto/md5"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	movieTitleRE = regexp.MustCompile(`^(.*?)(\([\d\?]{4}\/?[XLVI]*\))( *(?:\(TV\)|\(V\)|\(VG\))? *(?:\{.*\})?).*`)
)

type Title struct {
	movieId string
	Title   string
	Year    string
}

func (t *Title) MovieId() string {
	return t.movieId
}

/**
CRC: 0xE0CD55E1  File: movies.list  Date: Fri Mar 17 00:00:00 2017

Copyright 1991-2017 The Internet Movie Database Ltd. All rights reserved.

http://www.imdb.com

movies.list

17 Mar 2017

-----------------------------------------------------------------------------

MOVIES LIST
===========

"!Next?" (1994)                                         1994-1995
"#1 Single" (2006)                                      2006-????
"#1 Single" (2006) {Cats and Dogs (#1.4)}               2006
"#1 Single" (2006) {Finishing a Chapter (#1.5)}         2006
**/
func ExportTitlesDataset(inputPath, outputPath string) error {

	scanner, err := NewScanner(inputPath)
	if err != nil {
		return err
	}
	defer scanner.Close()

	scanner.Skip(15)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.HasPrefix(text, "----") {
			break
		}

		title, year, err := parseTitleAndYear(text)
		if err != nil {
			return err
		}

		fmt.Printf("%s %y\n", title, year)
	}

	return scanner.Err()
}

func parseTitleAndYear(expr string) (title string, year int, err error) {
	titleExpr := movieTitleRE.FindStringSubmatch(expr)
	if len(titleExpr) == 0 {
		err = fmt.Errorf("no match for the expression: `%s`", expr)
		return
	}

	normalizedTitle := strings.Trim(titleExpr[1], `"' `)

	title = normalizedTitle

	normalizedYear := strings.Trim(titleExpr[2], `(VLVI/?)`)

	if parsedYear, err := strconv.Atoi(normalizedYear); err == nil {
		year = parsedYear
	} else {
		year = -1
	}

	return
}

func makeId(title string, year int) string {
	id := []byte(fmt.Sprintf("%s;%d", title, year))
	return fmt.Sprintf("%x", md5.Sum(id))
}
