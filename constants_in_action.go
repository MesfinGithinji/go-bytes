/*
  The error here is cannot assign PI
  because i already assigned it as a constant
  so line 20 wont work
*/

package main

import "fmt"

func main() {
	const PI = 3.142
	const e = 2.72
	const gravity = 9.8

	fmt.Println(PI)
	fmt.Println(e)
	fmt.Println(gravity)

	PI := 20

}
