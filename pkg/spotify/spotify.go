package spotify

import (
	"fmt"

	"github.com/hamzaanis/go-spotify-account-generator/pkg/utils"
)

// CreateAccount creates a spotify account
func CreateAccount() {
	email := fmt.Sprintf("%v@gmail.com", utils.RandomString(15))
	password := utils.RandomString(15)

	flag := utils.MakeRequest(email, password)
	if flag {
		fmt.Printf("Spotify Account: %v, %v\n", email, password)
	}
}
