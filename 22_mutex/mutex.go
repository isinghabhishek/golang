package main

import (
	"fmt"
	"sync"
)

//  go run 22_mutex/mutex.go

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()

}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	// myPost.inc()
	// myPost.inc()
	wg.Wait()
	fmt.Println(myPost.views)
}
