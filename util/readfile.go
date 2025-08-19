package util

import (
	"bufio"
	"io"
	"os"
)

type AsciiLine []rune

type InmemFile struct {
	buf     []AsciiLine
	linePtr int
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

func (inf *InmemFile) HasNext() bool {
	return inf.linePtr < len(inf.buf)
}

func (inf *InmemFile) GetNext() AsciiLine {
	if inf.linePtr >= len(inf.buf) {
		return nil
	}

	line := inf.buf[inf.linePtr]
	inf.linePtr++

	return line
}
