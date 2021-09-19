package cmd

import (
	"time"

	"github.com/hamzaanis/go-spotify-account-generator/pkg/spotify"
)

// Starting point of the cmd
func Start() {
	for {
		spotify.CreateAccount()
		time.Sleep(1000)
	}
}
