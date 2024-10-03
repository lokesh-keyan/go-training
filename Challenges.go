package main

import (
	"fmt"
)

func Challenges() {
	//FizzBuzz()
	EvenEnded()
}

func FizzBuzz() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func EvenEnded() {
	count := 0
	for a := 1000; a < 9999; a++ {
		for b := a; b < 9999; b++ {
			n := a * b
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Printf("The number of even ended number: %d\n", count)
}
