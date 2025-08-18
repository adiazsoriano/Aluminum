package util

import (
	"bufio"
	"io"
	"os"
)

type InmemFile struct {
	buf [][]byte
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
		line, err := reader.ReadBytes(';')
		inf.buf = append(inf.buf, line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return inf, nil
}
