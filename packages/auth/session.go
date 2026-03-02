package auth

import "fmt"

func GetSession(username string) string {
	return fmt.Sprintf("Session for user: %s", username)
}
