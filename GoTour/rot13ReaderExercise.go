package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (wrap rot13Reader) Read(p []byte) (int, error) {
	length, err := wrap.r.Read(p)
	
	if err == io.EOF {
		return 0, io.EOF
	}
	
	for i := 0; i < length; i++ {
		if unicode.IsLetter(rune(p[i])) {
			if unicode.IsLetter(rune(p[i] + 13)) {
				p[i] = p[i] + 13
			} else {
				p[i] = p[i] + 13 - 26 // 25 -> alphabet length
			}
		}
		
	}
	
	return length, err
	
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
