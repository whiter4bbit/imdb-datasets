package db

import (
	"github.com/couchbase/vellum"
	"os"
)

type FSTWriter struct {
	f *os.File
	b *vellum.Builder
}

func NewFSTWriter(path string) (*FSTWriter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	b, err := vellum.New(f, nil)
	if err != nil {
		return nil, err
	}

	return &FSTWriter{
		f: f,
		b: b,
	}, nil
}

func (w *FSTWriter) Write(title []byte, id uint32) error {
	return w.b.Insert(title, uint64(id))
}

func (w *FSTWriter) Close() error {
	if err := w.b.Close(); err != nil {
		return err
	}
	return w.f.Close()
}
