package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/hamzaanis/go-spotify-account-generator/pkg/constants"
)

//RandomString generates a random string with the size n
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// MakeRequest makes and do the request with proper headers
func MakeRequest(email, password string) bool {
	payload := fmt.Sprintf("birth_day=3&birth_month=05&birth_year=1996&collect_personal_info=undefined&creation_flow=&creation_point=https://www.spotify.com/us/&displayname=BOT&gender=male&iagree=1&key=a1e486e2729f46d6bb368d6b2bcda326&platform=www&referrer=&send-email=0&thirdpartyemail=1&email=%v&password=%v&password_repeat=%v", email, password, password)
	req, _ := http.NewRequest("POST", constants.URL, bytes.NewBuffer([]byte(payload)))
	for k, v := range constants.HEADERS {
		req.Header.Set(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil && resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()
	return true
}
