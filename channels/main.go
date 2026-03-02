package main

import "fmt"

// sender → channel → receiver
// func processNum(numChan chan int) {

// 	for num := range numChan {

// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }
// func sum(result chan int, a, b int) {
// 	sumResult := a * b
// 	result <- sumResult
// }

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Task is running...")
	done <- true
}

func main() {
	done := make(chan bool)
	go task(done)

	<-done
	fmt.Println("Task completed.")
	// //1. Create a channel
	// numChan := make(chan int)

	// //2. Start a goroutine to process numbers from the channel
	// go processNum(numChan)

	// //3. Send random numbers to the channel
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// results := make(chan int)
	// go sum(results, 10, 20) //blocking call until the result is sent to the channel
	// fmt.Println(<-results)

}

// main → generates random number
// main → sends to channel
// processNum → receives
// processNum → sleeps 1 second
// repeat
