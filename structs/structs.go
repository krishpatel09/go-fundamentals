package main

import "fmt"

// order structs
// type Order struct {
// 	id        int
// 	amount    float32
// 	status    string
// 	createdAt time.Time //nanosecond precision
// }

// func (o *Order) changeStatus(status string) Order {
// 	o.status = status
// 	return *o
// }

func main() {

	person := struct {
		name string
		age  int
		city string
	}{
		"krish",
		25,
		"Delhi",
	}

	fmt.Printf("Person: %+v\n", person)

	// myOrder := Order{
	// 	id:     1,
	// 	amount: 100.50,
	// 	status: "pending",
	// }

	// myOrder.changeStatus("confirm")
	// fmt.Println(myOrder)

	// 	myOrder.createdAt = time.Now()
	// 	fmt.Println(myOrder.status)
	// 	fmt.Printf("Order: %+v\n", myOrder)

	// 	myOrder2 := Order{
	// 		id:        2,
	// 		amount:    200.75,
	// 		status:    "completed",
	// 		createdAt: time.Now(),
	// 	}
	// 	fmt.Printf("Order 2:", myOrder2)

}
