package main

import "fmt"

type gateway interface {
	payment(amount float32)
}

type payment struct {
	getway gateway
}

func (p payment) process(amount float32) {
	p.getway.payment(amount)
}

type razorpay struct{}

func (rp razorpay) payment(amount float32) {
	fmt.Println("make payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) payment(amount float32) {
	fmt.Println("make payment using stripe", amount)
}

func main() {

	// razor := razorpay{}
	// p1 := payment{getway: razor}
	// p1.process(100)

	// stripe := stripe{}
	// p2 := payment{getway: stripe}
	// p2.process(200)

	stripegw := stripe{}
	newPayment := payment{
		getway: stripegw}
	newPayment.process(300)
}
