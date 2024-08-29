package main

import "fmt"

type payment struct{}

func (p payment) makePayment(amount float32) {
	razorpayPaymentGateWay := razorpay{}
	razorpayPaymentGateWay.pay(amount)
}

// go run 15_interface/interface.go

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("makeing using razorpay", amount)

}

func main() {
	newPayment := payment{}
	newPayment.makePayment(100)
}
