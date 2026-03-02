package main

import (
	"fmt"
	"time"
)

func emailSender(emailchannel chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emailchannel {
		fmt.Println("Sending email:", email)
		time.Sleep(time.Second)

	}
	<-done
}

func main() {

	emailchannel := make(chan string, 2) //buffered channel
	done := make(chan bool)              //unbuffered channel
	go emailSender(emailchannel, done)

	for i := 0; i < 5; i++ {
		emailchannel <- fmt.Sprintf("%d@gmail.com", i)
	}
	close(emailchannel)
	<-done
	fmt.Println("All emails processed successfully...")
}
