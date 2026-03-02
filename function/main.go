package main

// func add(a int, b int) int {
// 	return a + b
// }

// func language() (string, string, string) {
// 	return "Go", "javascript", "python"
// }

func processIt() func(a int) int {
	return func(a int) int {
		return a * 2
	}
}

func main() {
	// result := add(3, 5)
	// println("The result is:", result)
	// lang1, lang2, lang3 := language()
	// println("Languages:", lang1, lang2, lang3)

	doubleFunc := processIt()
	result := doubleFunc(4)
	println("The result is:", result)
}
