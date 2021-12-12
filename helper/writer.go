package helper

import (
	"encoding/csv"
	"os"
)

type Writer struct {
	f *os.File
}

func NewWriter(filename string) (*Writer, error) {
	err := os.Remove(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &Writer{
		f: f,
	}, nil
}

func (w *Writer) Write(records [][]string) error {
	wr := csv.NewWriter(w.f)
	err := wr.WriteAll(records)
	if err != nil {
		return err
	}
	wr.Flush()

	err = wr.Error()
	if err != nil {
		return err
	}

	return nil
}

func (w *Writer) Close() error {
	return w.f.Close()
}
