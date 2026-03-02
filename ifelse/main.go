package main

import (
	"fmt"
)

func main() {

	// age := 15

	// IF-ELSE
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else if age >= 13 {
	// 	fmt.Println("You are a teenager.")
	// } else {
	// 	fmt.Println("You are a child.")
	// }

	// SWITCH (without variable)
	// switch {
	// case age >= 18:
	// 	fmt.Println("You are an adult.")
	// case age >= 13:
	// 	fmt.Println("You are a teenager.")
	// default:
	// 	fmt.Println("You are a child.")
	// }

	// // SWITCH with time
	// switch time.Now().Weekday() {
	// case time.Monday, time.Tuesday:
	// 	fmt.Println("It's work day.")
	// case time.Wednesday, time.Thursday, time.Friday:
	// 	fmt.Println("It's midweek.")
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend.")
	// default:
	// 	fmt.Println("It's regular day.")
	// }

	// Type checking function calls
	typecheck(23.7)
	// typecheck("hello")
	// typecheck()
}

// Function declaration (outside main)
func typecheck(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("i is an int: %d\n", t)
	case string:
		fmt.Printf("i is a string: %s\n", t)
	default:
		fmt.Printf("i is of unknown type: %T\n", t)
	}
}
