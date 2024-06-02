package main

import (
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, io.EOF
	}
	
	for i, _ := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
