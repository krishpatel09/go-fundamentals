package main

import "fmt"

type Status int

const (
	Pending Status = iota
	Confirmed
	Shipped
	Delivered
)

func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Confirmed:
		return "Confirmed"
	case Shipped:
		return "Shipped"
	case Delivered:
		return "Delivered"
	default:
		return "Unknown"
	}
}

func main() {
	var orderStatus Status = Delivered
	fmt.Println(orderStatus)
}
