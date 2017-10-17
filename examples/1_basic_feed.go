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
	fb.SetAccessToken(`EAACEdEose0cBAJ4cTiwJkXbHMZAaBq5R4JK33IedxhWe5g5Sn0OWVQz85cYb1RSuV2idZAoGbtGPSZAWZBqApQ71MbigYDptl0o7li1luYIcgae5Q8ZC00qrNZAsKATFMS7AIEsgEzbYQMZAt58e2zVzoR0brphfm67QNkZBw3lCrHjDpM0lJD0Uvf2EHwlcsTZBlyUQQZCFYZC8wZDZD`)

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
