package main

import (
	"fmt"
	"maps"
)

// maps --> hash, object, dict
func main() {
	// creating map
	// m := make(map[string]string) // go run 10_maps/maps.go

	// // setting an element
	// m["name"] = "golang"
	// m["area"] = "backend"

	// getting an element
	// fmt.Println(m["name"], m["area"])
	// IMP: If key does not exists in the map then it returns zero

	// m := make(map[string]int) // go run 10_maps/maps.go
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])
	// fmt.Println(len(m))

	// // delete(m, "price")
	// clear(m)
	// fmt.Println(m)

	// map without make
	// m := map[string]int{"price": 358, "phone": 3}

	// fmt.Println(m)

	m1 := map[string]int{"price": 358, "phone": 3}
	m2 := map[string]int{"price": 3589, "phone": 4}

	fmt.Println(maps.Equal(m1, m2))

}
