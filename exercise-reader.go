// Mettre en place un type Reader qui émet un flux infini de caractères ASCII 'A'.

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (MyReader) Read(b []byte) (int, error) {
	bs := []byte("A")

	for i, v := range bs {
		b[i] = v
	}

	return len(bs), nil
}

func main() {
	reader.Validate(MyReader{})
}
