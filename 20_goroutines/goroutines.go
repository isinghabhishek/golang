package main

import (
	"fmt"
	"sync"
	"time"
)

// go run 20_goroutines/goroutines.go

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		// go task(i)
		wg.Add(1)
		go task(i, &wg)

		wg.Wait()
		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}
	time.Sleep(time.Second * 2)
}
