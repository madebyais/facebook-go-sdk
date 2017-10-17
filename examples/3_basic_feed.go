package main

import (
	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)

// BasicFeed represents the basic of
// how to use facebook-go-sdk
func BasicFeed() {
	// initalize facebook-go-sdk
	fb := facebook.New()

	// set your access token
	// NOTES: Please exchange with your access token
	fb.SetAccessToken(`...`)

	// and directly get your feed :)
	data, err := fb.API(`/me/feed`).Get()
	if err != nil {
		panic(err)
	}

	// print your feed
	fmt.Println(`
		
		## SAMPLE - FEED
	`)
	fmt.Println(data)
}
