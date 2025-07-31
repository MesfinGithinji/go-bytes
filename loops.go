package main

import "fmt"

func main() {

	nums := []int{10, 21, 30, 43, 50}

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			fmt.Println(nums[i])
		} else {
			fmt.Printf("The number at %d is %d and it is odd\n", i, nums[i])
		}
	}
}
