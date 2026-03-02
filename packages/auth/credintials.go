package auth

import (
	"fmt"

	"github.com/KrishPatel/go-fundamentals/user"
)

func LoginWithCredentials(username string, password string) {

	u := user.User{
		Username: "krish",
		Password: "krish@123",
	}

	if u.Username == username && u.Password == password {
		fmt.Println("Login successful!")
		fmt.Println(GetSession(username))
	} else {
		fmt.Println("Login failed!")
	}
}
