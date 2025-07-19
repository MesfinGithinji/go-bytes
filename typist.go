/*
Type inference challenge
An integer
A float
A string
A boolean
Print each variable along with its Go-inferred type using fmt.Printf("Type: %T\n", var).
Also look at the various format specifiers
*/

package main

import "fmt"

func main() {

	age := 25
	name := "Mesh"
	weight := 68.59
	feel := true

	fmt.Printf("The value %d is of Type: %T at position %p in memory\n", age, age, &age) //playing around with pointers here
	fmt.Printf("The value %s is of Type: %T\n", name, name)
	fmt.Printf("The value %f is of Type: %T\n", weight, weight)
	fmt.Printf("The value %t is of Type: %T\n", feel, feel)

	msg := fmt.Sprintf("Hello, %s. You are %d years old.", name, age) //how to build a string message
	fmt.Println(msg)

}
