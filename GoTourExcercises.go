package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for z < 10.0 {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func loopsAndFunctions() {
	fmt.Println(Sqrt(2))
}
