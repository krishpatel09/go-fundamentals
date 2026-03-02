package main

import (
	"fmt"
	"sync"
)

// type post struct {
// 	views int
// 	mu    sync.Mutex
// }

// func (p *post) inc(wg *sync.WaitGroup) {
// 	defer func() {
// 		p.mu.Unlock()
// 		wg.Done()
// 	}()

// 	// p.mu.Lock()
// 	p.mu.Lock()
// 	p.views += 1
// 	// p.mu.Unlock()
// }

type Counter struct {
	val int
	mu  sync.Mutex
}

func (c *Counter) Inc() {
	c.mu.Lock() //don't use main fuction
	c.val++
	c.mu.Unlock()
}

func main() {
	// var wg sync.WaitGroup
	// myPost := post{views: 0}

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go myPost.inc(&wg)

	// }

	// wg.Wait()
	// fmt.Println(myPost.views)

	var wg sync.WaitGroup
	Counter := Counter{}

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(Counter.val)
}
