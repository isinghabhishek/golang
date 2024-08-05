package main

import "fmt"

// iterating over data structure
// go run 11_Range/range.go
func main() {
	// nums := []int{1, 1, 3, 5, 8, 13}
	// printing array
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// sum of array
	// sum := 0
	// //  indx
	// for _, num := range nums {
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	// map
	// m := map[string]string{"fname": "john", "lname": "doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// 	fmt.Println(k, v)
	// }

	// for k := range m {
	// 	fmt.Println(k)
	// }

	// unicode code point rune
	// starting byte of rune
	// 255 --> 1 byte, 2 byte
	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}
