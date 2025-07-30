package main

import "fmt"

func main() {

	var age int = 30
	var p *int = &age

	fmt.Printf("The variable is age and the value is %d. \nThe pointer is p and the addrss is %p.\n", age, p)
}
