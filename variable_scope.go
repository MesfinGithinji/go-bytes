/*
 Variables Scope Experiment
*/

package main

import "fmt"

func main() {

	myName := "Mesh"
	fmt.Println(myName)
	/*
	  this will print Mesh because Go will find myname
	  declared under the scope of the main function and thus myName is scoped only to the main function
	  while the variable below this , although same variable name , belongs to a different scope thus
	  will print Rich.
	  The two variables although named the same , are scoped differently.
	  Variable in Go follow block-level scope
	*/

	{
		var myName string = "Rich"
		fmt.Println(myName) //this will print Rich
	}
}
