package main

/*
 Use variables to:

    Take two numbers (you can hardcode them).

    Perform: addition, subtraction, multiplication, and division.

    Print results with formatted strings like "Sum of %d and %d is %d"
*/

import "fmt"

func main() {
	//declare my two variables
	num1 := 100
	num2 := 40

	//addition
	sum := num1 + num2
	result := fmt.Sprintf("The Sum of %d and %d is = %d", num1, num2, sum)
	fmt.Println(result)

	//subtraction
	diff := num1 - num2
	result2 := fmt.Sprintf("The Difference of %d and %d is = %d", num1, num2, diff)
	fmt.Println(result2)

	//multiplication
	prod := num1 * num2
	result3 := fmt.Sprintf("The Product of %d and %d is = %d", num1, num2, prod)
	fmt.Println(result3)

	//division
	div := float64(num1) / float64(num2)
	result4 := fmt.Sprintf("%d divide by %d is = %.2f", num1, num2, div)
	fmt.Println(result4)

	/*
	 This would also work
	 num1, num2 := 100, 40

	 fmt.Printf("Sum of %d and %d is %d\n", num1, num2, num1+num2)
	 fmt.Printf("Difference of %d and %d is %d\n", num1, num2, num1-num2)
	 fmt.Printf("Product of %d and %d is %d\n", num1, num2, num1*num2)
	 fmt.Printf("Quotient of %d divided by %d is %d\n", num1, num2, num1/num2)
	 could also convert to float64 for better accuracy
	*/
}
