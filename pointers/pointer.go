package main

import "fmt"

// func changenum(num int) {
// 	num = 20
// 	fmt.Println("change num", num)
// }

//by reference
func changenum(num *int) {
	*num = 20
	fmt.Println("change num", *num)
}

func main() {
	num := 10

	// fmt.Println("memory address of num:", &num)
	changenum(&num)
	fmt.Println("main num", num)
}
