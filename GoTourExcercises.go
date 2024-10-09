package main

import (
	"fmt"
	"math"
	"strings"
)

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

func slicesOfSlices(){
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i:=0; i < len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i], " "));
	}
}

func appendSlice(){
	var s []int
	printslice(s)

	s = append(s, 0)
	printslice(s)

	s = append(s, 2,3,4)
	printslice(s)
}

func printslice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func rangeSlice(){
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow{
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i := range pow{
		fmt.Printf("2**%d = %d\n", i, pow[i])
	}

	for _, value := range pow{
		fmt.Printf("%d\n", value)
	}

	for i, _ := range pow{
		fmt.Println(pow[i])
	}
}

// func Pic(dx, dy int) [][]uint8 {
//     result := make([][]uint8, dy)
//     for y := 0; y < dy; y++ {
//         result[y] = make([]uint8, dx)
//         for x := 0; x < dx; x++ {
//             result[y][x] = uint8((x + y) / 2)
//         }
//     }
//     return result
// }

type Vertex struct{
	Lat, Long float64
}

func mapsLiterals(){
	var m = map[string]Vertex{
		"Atlanta": {
			40.68433, -74.39967,
		},
		"Chennai": {
			37.42202, -122.08408,
		},
	}

	a := m["Atlanta"]
	v, k := m["Atlanta"]

	fmt.Println(m)
	fmt.Println(a.Lat)
	fmt.Println("The value:", v, "Present?", k)
}

func WordCount(s string){
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words{
		if v, k := m[word]; k{
			m[word] = v + 1
		} else{
			m[word] = 1
		}
	}

	fmt.Println(m)
}

func compute(fn func(float64, float64) float64) float64{
	return fn(3, 4)
}

func functionValues(){
	hypot := func(x, y float64) float64{
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func funcClosures() {
	pos, neg := adder(), adder() // pos and neg are values returned by the adder which is a func 
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func(int) int {
	a := 0;
	b := 1
	return func(x int) int{
		c := a + b
		a = b
		b = c
		return c
	}
}

func funcClosureExcercise() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}