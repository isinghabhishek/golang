package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

// go run 14_closures/closures.go

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
}
