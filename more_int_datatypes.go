package main

import "fmt"

func main() {
	//create a variable ofof type uint8- positive 8 bit value range is 0-255
	var mySmallPositive uint8 = 8
	fmt.Printf("My value is :%d  and its of type %T\n", mySmallPositive, mySmallPositive)

	//variable of type int8 - either positive or negative range -128 - 127.
	var mixedNum int8 = -120
	fmt.Printf("The value is %d and the type is %T\n", mixedNum, mixedNum)

	var myNewInt = string(mySmallPositive) //type casting can only be done for types that can be converted
	fmt.Printf("New Type of mySmallPositive is %T\n", myNewInt)

	//byte data type - alias for uint8 - raw data value
	//eg this example below will print the ASCII value of A which is 65
	var a byte = 'A'
	fmt.Println(a)

	//rune data type - alias for int32
	//used to represent Unicode code points — essentially, a single character.
	//It's how Go handles characters in a string at the individual level.
	var b rune = 'B'
	fmt.Println(b)
}
