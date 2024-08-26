package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

//  go run 16_structs/struct.go

func main() {
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

	myOrder.createdAt = time.Now()
	fmt.Println("Order struct", myOrder)
}
