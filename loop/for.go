package main

import (
	"fmt"
)

func main() {
	//while loop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i = i + 1

	// }

	//loop for classis
	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//range loop
	for i := range 3 {
		fmt.Println(i)
	}
}
