package datasets

import (
	"bufio"
	"compress/gzip"
	"os"
)

type Scanner struct {
	s *bufio.Scanner
	f *os.File
}

func NewScanner(path string) (*Scanner, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	return &Scanner{
		s: bufio.NewScanner(gzipReader),
		f: file,
	}, nil
}

func (s *Scanner) Skip(n int) {
	for n > 0 && s.s.Scan() {
		n -= 1
	}
}

func (s *Scanner) Scan() bool {
	return s.s.Scan()
}

func (s *Scanner) Text() string {
	return s.s.Text()
}

func (s *Scanner) Err() error {
	return s.s.Err()
}

func (s *Scanner) Close() {
	s.f.Close()
}
