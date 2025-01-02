package main

import (
	"fmt"
	"go-training/leetcode"
	"go-training/goTour"
)

func main() {
	fmt.Println("Go Training Repository")
	fmt.Println("Select a category:")
	fmt.Println("1. LeetCode")
	fmt.Println("2. GoTour")

	var choice int
	// Get user input for choice
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// switch choice {
	// case 1:
	// 	leetcode.Main()
	// case 2:
	// 	goTour.Main()
	// default:
	// 	fmt.Println("Invalid choice")
	// }
}
