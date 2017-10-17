package main

import (
	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)

// BasicFeedDelete will show you about
// how to delete a feed from your timeline
func BasicFeedDelete() {
	// initalize facebook-go-sdk
	fb := facebook.New()

	// set your access token
	// NOTES: Please exchange with your access token
	fb.SetAccessToken(`...`)

	// delete a feed by submitting feed_id
	// NOTES: Please exchange with your feed_id
	data, err := fb.API(`/1665362379_10213149191998820`).Delete()
	if err != nil {
		panic(err)
	}

	fmt.Println(`

		## SAMPLE - DELETE A FEED
	`)

	fmt.Println(data)
}
