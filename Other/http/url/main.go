package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://www.googleapis.com/youtube/v3/playlists")
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("part", "snippet,contentDetails")
	q.Set("mine", "true")
	q.Set("maxResults", "25")

	u.RawQuery = q.Encode()
	fmt.Println(u.String())
}
