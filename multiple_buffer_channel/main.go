package main

func main() {

	chan1 := make(chan int, 2)    //buffered channel
	chan2 := make(chan string, 2) 

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Golang"
	}()

	for i := 0; i < 2; i++ {
		select {
		case val := <-chan1:
			println("Received from chan1:", val)
		case val := <-chan2:
			println("Received from chan2:", val)
		}
	}
}
