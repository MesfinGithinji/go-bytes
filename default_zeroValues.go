/*
* Lets learn about Go's zero values
 */

package main

import "fmt"

func main() {
	var i int
	var f float64
	var s string
	var b bool

	fmt.Println(i) //i is an int zero value is 0
	fmt.Println(f) //f is a float zero value is 0
	fmt.Println(s) //s is a string zero value is an empty string
	fmt.Println(b) //b is a boolean zero value is false

}
