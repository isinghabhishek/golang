package main

import (
	"fmt"
	"time"
)

// go run 21_channels/channels.go

// Send
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("processing...")
// }

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received date from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received date from chan1", chan2Val)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@singh.com", i)
	// }

	// fmt.Println("done sending...")
	// // this is imported
	// close(emailChan)
	// <-done
	// emailChan <- "abhishek@singh.com"
	// emailChan <- "abhi@singh.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)
	// go task(done)
	// <- done // block

	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println(res)

	// numChan := make(chan int)
	// go processNum(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }
	// messageChan <- "ping" // blocking
	// msg := <-messageChan
	// fmt.Println(numChan)
}
