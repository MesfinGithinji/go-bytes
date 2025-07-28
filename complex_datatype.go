package main

import "fmt"

func main() {

	//be careful when assigning float32 or float64
	//the memory is different
	var smallFloat float32
	smallFloat = 23.02
	fmt.Println(smallFloat)

	//complex is a combination of real and imaginary number
	//eg 3 + 4i
	//4i is imgainary
	var myComplex complex128 = 345
	fmt.Println(myComplex)

	fmt.Println()
	fmt.Println(real(myComplex))
	fmt.Println(imag(myComplex))

	/*
		  Why Use Complex Numbers?
		   They’re used in:
			Signal processing 🎧
			Physics 🌌
			Math and engineering 🧮
	*/
}
