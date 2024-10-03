package main

import "fmt"

func main() {
	//
	//FizzBuzz()
	//strings()
	//EvenEnded()
	loopsAndFunctions()
}

func types() {
	x, y := 10.0, 0.0

	fmt.Printf("x=%v\n", x)
	fmt.Printf("y=%v\n", x)

	if mean := (x + y) / 2.0; mean > 5 {
		fmt.Printf("result: %v, type of %T\n", mean, mean)
	} else {
		fmt.Printf("bad result: %v, type of %T\n", mean, mean)
	}
}

func conditions(mean float64) {
	if mean > 5 {
		fmt.Println("x is big")
	}
}

func loop() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
