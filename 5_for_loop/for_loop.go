package main

import "fmt"

// for --> only construct in go for looing
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
