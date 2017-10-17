package main

import (
	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)

// BasicMe represents the basic of
// how to use facebook-go-sdk
func BasicMe() {
	// initalize facebook-go-sdk
	fb := facebook.New()

	// set your access token
	// NOTES: Please exchange with your access token
	fb.SetAccessToken(`...`)

	// and directly get your feed :)
	data, err := fb.API(`/me`).Get()
	if err != nil {
		panic(err)
	}

	// print your feed
	fmt.Println(`
    
		## SAMPLE - ME
	`)
	fmt.Println(data)
}
