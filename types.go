//lets play around with basic data types in Go
//lets also understand variable declaration and initialization
//var deaclares one or more variables

// := is shorthand for declaring and initializing a variable only inside a function.
//only inside a function and this has type inference

// use const for constant variables whose values do not change

package main

import "fmt"

func main() {

	//var can be used for many variable declarations at once
	var a, b, c int
	a = 5
	b = 15
	c = 20
	fmt.Println(a, b, c)

	//boolean is either true or false
	//zero value for bool is false
	var myBool bool
	myBool = true
	fmt.Println(myBool)

	//strings
	var myString string
	myString = "This is a String"
	fmt.Println(myString)

	//number we have integers , float and complex numbers
	//zero value for integers is 0
	var myInt int
	myInt = 10
	fmt.Println(myInt)

	//constant
	const PI = 3.142
	fmt.Println(PI)

	//float eithe float64 or float32
	var myfloat float64
	myfloat = 98.78
	fmt.Println(myfloat)
}
