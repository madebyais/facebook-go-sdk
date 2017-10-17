package main

import (
	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)

// BasicFeedPost will show you about
// how to post a feed to your timeline
func BasicFeedPost() {
	// initalize facebook-go-sdk
	fb := facebook.New()

	// set your access token
	// NOTES: Please exchange with your access token
	fb.SetAccessToken(`EAACEdEose0cBAJ4cTiwJkXbHMZAaBq5R4JK33IedxhWe5g5Sn0OWVQz85cYb1RSuV2idZAoGbtGPSZAWZBqApQ71MbigYDptl0o7li1luYIcgae5Q8ZC00qrNZAsKATFMS7AIEsgEzbYQMZAt58e2zVzoR0brphfm67QNkZBw3lCrHjDpM0lJD0Uvf2EHwlcsTZBlyUQQZCFYZC8wZDZD`)

	// submit your feed
	data, err := fb.API(`/me/feed`).Messages(`coba yo`).Post()
	if err != nil {
		panic(err)
	}

	fmt.Println(`

		## SAMPLE - POST A FEED
	`)

	fmt.Println(data)
}
