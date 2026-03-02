package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Printf("Task %d is running\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		task(i, &wg)

	}
	wg.Wait()
}
