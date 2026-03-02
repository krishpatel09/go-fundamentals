package main

//  variadic function is one that accepts a variable number of arguments of the same type

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	println(sum(nums...))
	println(sum(1, 2, 3))
	println(sum(4, 5, 6, 7))
}
