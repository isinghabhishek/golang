package main

import "fmt"

// numbered sequence of specific length
func main() {

	// 1. Int array
	// at the time array intialisation it get store with zero value
	// zeroed values
	// int -> 0, string -> "", bool --> false
	// var nums [4]int
	// nums[0] = 3
	// fmt.Println(nums[0])
	// fmt.Println(nums)
	// // length of array
	// fmt.Println(len(nums)) // go run 8_arrays/arrays.go

	// 2. boolean array
	// go run 8_arrays/arrays.go
	// var nums [4]bool
	// nums[0] = true
	// fmt.Println(nums)

	// //3.String
	// var name [4]string
	// name[0] = "Abhi"
	// fmt.Println(name)

	// declrating value in single line
	// nums := [3]int{1, 3, 5}
	// fmt.Println(nums)

	// 2d array
	nums := [2][2]int{{3, 5}, {8, 11}}
	fmt.Println(nums)

	// - fixed size, that is predictable
	// - Memory optimazation
	// - Constant time access
}
