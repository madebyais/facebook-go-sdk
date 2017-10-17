Facebook Go SDK
------------------------
###### v0.0.1

Facebook Go SDK is an unofficial Facebook SDK for Golang.
It is very simple and easy to use.

### SAMPLE

```
package main

import (
	"fmt"

	facebook "github.com/madebyais/facebook-go-sdk"
)

func main() {
	// initalize facebook-go-sdk
	fb := facebook.New()

	// set your access token
	// NOTES: Please exchange with your access token
	fb.SetAccessToken(`xxxxxx`)

	// and directly get your feed :)
	data, err := fb.API(`/me/feed`).Get()
	if err != nil {
		panic(err)
	}

	// print your feed
	fmt.Println(data)
}

```

Please refer to `examples` folder for more details.

### NEXT

- Enable middelware handler to retrieve access_token

### AUTHOR

[Faris](madebyais@gmail.com)

### LICENSE

Release under [MIT LICENSE](https://github.com/madebyais/facebook-go-sdk/blob/master/LICENSE)
