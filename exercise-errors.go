package main

import (
	"fmt"
	"strconv"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		e := ErrNegativeSqrt(x)
		return 0, e
	}

	z := float64(1)
	var zn float64

	for i := 0; i < 1000; i++ {
		if zn != 0 {
			z = zn
		}

		zn = z - (((z * z) - x) / (2 * z))
	}

	return zn, nil
}

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + strconv.FormatFloat(float64(e), 'f', -1, 64)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
