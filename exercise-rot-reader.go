// Mettre en place un rot13Reader qui implémente io.Reader et lit un io.Reader,
// modifiant le flux en appliquant la substitution rot13 de chiffrement à tous les caractères alphabétiques.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r2 rot13Reader) Read(b []byte) (int, error) {
	n, err := r2.r.Read(b)

	if err != nil {
		return n, err
	}

	for i, v := range b[:n] {
		switch {
		case v >= []byte("A")[0] && v <= []byte("Z")[0]:
			if (v + 13) <= []byte("Z")[0] {
				b[i] = v + 13
			} else {
				b[i] = v + 13 - []byte("Z")[0] + []byte("A")[0] - 1
			}

		case v >= []byte("a")[0] && v <= []byte("z")[0]:
			if (v + 13) <= []byte("z")[0] {
				b[i] = v + 13
			} else {
				b[i] = (v + 13) - []byte("z")[0] + []byte("a")[0] - 1
			}
		default:
			// caractère non alphabétiques
		}
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
