package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age >= 0 && age <= 12 {
		fmt.Println("You are a child.")
	} else if age >= 13 && age <= 19 {
		fmt.Println("You are a teen.")
	} else if age >= 20 && age <= 64 {
		fmt.Println("You are an adult.")
	} else if age >= 65 {
		fmt.Println("You are a senior.")
	} else {
		fmt.Println("Invalid age entered.")
	}
}
