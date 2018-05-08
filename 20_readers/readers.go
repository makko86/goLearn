package main

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/tour/reader"
)

//populates slice until EOF
func readerExample() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v | err = %v | b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

//Exercise Reader that emits 'A'
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	size := len(b)
	for i := 0; i < size; i++ {
		b[i] = 'A'
	}
	return size, nil
}

//Exercise 2: ROT-13 Reader
type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(buf []byte) (int, error) {
	n, err := r.r.Read(buf)
	if err != nil {
		for i := 0; i < n; i++ {
			buf[i] = rot13(buf[i])
		}
	}
	return n, err
}

func rot13(b byte) byte {
	//TODO...
	return b
}

func main() {
	readerExample()
	reader.Validate(MyReader{})
}
