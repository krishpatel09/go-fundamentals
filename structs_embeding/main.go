package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
	city string
}

type Order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time
	User      //embedding
}

func main() {
	// newcustomer := User{
	// 	name: "krish",
	// 	age:  25,
	// 	city: "Rajkot",
	// }
	neworder := Order{
		id:        1,
		amount:    100.50,
		status:    "pending",
		createdAt: time.Now(),
		User: User{
			name: "krish",
			age:  25,
			city: "Rajkot",
		},
	}

	fmt.Println("new order:", neworder)
}
