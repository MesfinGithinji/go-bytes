package main

import "fmt"

func main() {

	//strings in Go lang use double quotes
	//zero value in strings is "" an empty string
	var myName string = "Mesfin"
	fmt.Println(myName)

	//we can also use back ticks
	var secondName string = `mezi`
	fmt.Println(secondName)

	//we can use formatters like Printf to format strings
	//Printf  formats a string according to a format specifier then prints the string to std output
	var tool, subject string = "Terraform ", "Code"
	fmt.Printf("%s is used for Infrastructure as %s\n", tool, subject)

	//Sprintf also formats according to a format sepcifier but returns the formatted string
	msg := fmt.Sprintf("%s is used for Infrastructure as %s\n", tool, subject)
	fmt.Println(msg)

}
