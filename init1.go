/*
This is a GO program that
Declares 3 int variables and assigns them values in a single line
Declare float64, string ,bool using :=
print all values in separate lines

*/

package main

import "fmt"

func main() {
	/*
	 Since we are in a function , we can do this
	 a, b, c := 10, 20, 30
	 Go encourages minimal use of var in a function
	**/
	var a, b, c = 10, 20, 30

	myWeight := 63.45 //infers float
	myName := "mesh"  //infers sting
	mybool := true    //infers boolean

	fmt.Println(a, b, c)
	fmt.Println(myWeight)
	fmt.Println(myName)
	fmt.Print(mybool)
}
