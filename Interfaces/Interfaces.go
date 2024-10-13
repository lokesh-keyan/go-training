package main

import (
	"fmt"
	"time"
	//"math"
)

type I interface{
	M()
}

type T struct {
	S string
}

func (t *T) M(){
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M(){
	fmt.Println(f)
}

func main(){
	// interface nil
	// var i I 

	// var t *T
	// i = t
	// describe(i)
	// i.M()

	// i = &T{"hello"}
	// describe(i)
	// i.M()

	//empty interface
	// var j interface{}
	// describe(j)

	// j = 42
	// describe(j)

	// j = "hello"
	// describe(j)

	// // Type Assertions
	// var i interface{} = "hello"
	// s := i.(string)
	// fmt.Println(s)

	// s, ok := i.(string)
	// fmt.Println(s, ok)

	// f, ok := i.(float64)
	// fmt.Println(f, ok)

	// this will error out but not the top one
	// f = i.(float64) // panic
	// fmt.Println(f)

	//Stringer
	// hosts := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for name, ip := range hosts {
	// 	fmt.Printf("%s: %s\n", name, ip)
	// }

	//error interface
	err := &MyError{
		time.Now(),
		"it didn't work",
	}

	if err != nil {
		fmt.Println(err)
	}
}

func (ip IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

type IPAddr [4]byte

type MyError struct{
	When time.Time
	What string
}

func (e * MyError) Error() string{
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// this is equal to error = &MyError{}
// func run() error{

// }

// func describe(i I){
// 	fmt.Printf("(%v, %T)\n", i, i)
// }

func describe(j interface{}){
	fmt.Printf("(%v, %T)\n", j, j)
}

