package util

import (
	"bufio"
	"io"
	"iter"
	"os"
)

type AsciiLine []rune

type InmemFile struct {
	buf []AsciiLine
}

func NewInmemFile(inpFilePath string) (*InmemFile, error) {
	f, err := os.Open(inpFilePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	inf := &InmemFile{}

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadBytes('\n')
		inf.buf = append(inf.buf, AsciiLine([]rune(string(line))))
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return inf, nil
}

func (al AsciiLine) ColsRunes() iter.Seq2[int, rune] {
	return func(yield func(int, rune) bool) {
		for i, r := range al {
			if !yield(i, r) {
				return
			}
		}
	}
}

func (inf *InmemFile) RowsLines() iter.Seq2[int, AsciiLine] {
	return func(yield func(int, AsciiLine) bool) {
		for i, l := range inf.buf {
			if !yield(i, l) {
				return
			}
		}
	}
}
