package main

func main() {
	// nums := []int{6, 7, 8, 9}

	// for i, num := range nums {
	// 	println(i, num)
	// }

	m := map[string]string{
		"name": "krish",
		"age":  "25",
		"city": "ahmedabad",
	}

	for key, value := range m {
		println(key, value)
	}
}
