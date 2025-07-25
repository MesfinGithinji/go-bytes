package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Print("Enter your score (0-100): ")
	fmt.Scan(&score)

	if score < 0 || score > 100 {
		fmt.Println("Invalid score.")
	} else if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}
