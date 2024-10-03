package main

import (
	"fmt"
)

func strings() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))
	fmt.Printf("book[0] = %c (type %T)\n", book[0], book[0])
	fmt.Println(book[4:12]) // index 4 to 11 total 12 - 4 length
	fmt.Println(book[1:])   //from index 1 to end
	fmt.Println(book[:5])   // total length of 5 so index 0 to 4

	// string s are unicode
	// with back tick sign you can do multiple lines
}
