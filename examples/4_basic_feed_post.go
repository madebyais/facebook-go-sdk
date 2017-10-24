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
	fb.SetAccessToken(`...`)

	// submit your feed
	data, err := fb.API(`/me/feed`).Messages(`coba yo`).Post()
	if err != nil {
		panic(err)
	}

	fmt.Println(`

		## SAMPLE - POST A FEED
	`)

	fmt.Println(data)

	// if you want to share a link you can use the `Link` method
	// e.g. fb.API(`/me/feed`).Messages(`coba yo`).Link(`http://madebyais.com`).Post()
}
