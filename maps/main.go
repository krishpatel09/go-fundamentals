package main

import "fmt"

//maps => hase,object, dist
func main() {

	m := make(map[string]string)

	m["name"] = "krish" //set value
	m["age"] = "25"
	fmt.Println(m["name"], m["age"]) //get value
	delete(m, "age")                 //delete value
	fmt.Println(m["name"], m["age"]) //get value
	clear(m)
	fmt.Println(m["name"], m["age"]) //get value

}
