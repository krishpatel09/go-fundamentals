package main

import (
	"fmt"

	"github.com/KrishPatel/go-fundamentals/auth"
)

func main() {

	var username string
	var password string

	fmt.Print("Enter Username:")
	fmt.Scan(&username)

	fmt.Print("Enter Password:")
	fmt.Scan(&password)

	auth.LoginWithCredentials(username, password)

}
