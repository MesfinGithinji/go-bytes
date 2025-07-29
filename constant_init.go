package main

import "fmt"

func main() {

	//constans in Go can be either types or untyped
	//constants can be boolean
	const (
		mybool = true
		name   = "mesh" //can be string
		uno    = 'm'    //can be rune
		state  = 1 < 2  //an expression that can be evaluated at compiletime
	)

	fmt.Println(state)
}
