package main

import (
	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)

// GenerateAccessToken is a sample that shows you
// how to generate facebook access_token from server side
func GenerateAccessToken() {
	// initalize facebook-go-sdk
	fb := facebook.New()

	// set your application id (client_id)
	fb.SetAppID(`{APP_ID}`)

	// set your application secret (client_secret)
	fb.SetAppSecret(`{APP_SECRET}`)

	// Call https://www.facebook.com/v2.10/dialog/oauth?client_id={APP_ID}&redirect_uri={YOUR_REDIRECT_URI}
	// And then, retrieve the `code` value from {YOUR_REDIRECT_URI}?code=....

	// Then call GenerateAccessToken(redirectURI string, code string) method
	data, err := fb.GenerateAccessToken(`{YOUR_REDIRECT_URI}`, `{GENERATED_CODE}`)
	if err != nil {
		panic(err)
	}

	// print data object
	fmt.Printf(`%+v`, data)
}
