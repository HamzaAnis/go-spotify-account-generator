package cmd

import (
	"flag"
	"sync"

	"github.com/HamzaAnis/go-spotify-account-generator/pkg/spotify"
)

// Starting point of the cmd
func Start() {
	routines := flag.Int("routines", 10, "The number of go routines to run")
	size := flag.Int("size", 100, "The total number of accounts to generate")
	flag.Parse()

	chunk := *size / *routines
	var wg sync.WaitGroup
	for i := 0; i < *routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < chunk; i++ {
				spotify.CreateAccount()
			}
		}()
	}
	wg.Wait()
}
