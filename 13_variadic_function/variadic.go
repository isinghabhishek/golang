package main

import "fmt"

// go build 13_variadic/variadic.go

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}
	return total
}
func main() {
	nums := []int(1, 3, 5, 8)
	result := sum(nums...)
	fmt.Println(result)
	// fmt.Println(1, 2, 3, 4, 5, "hello")
}
