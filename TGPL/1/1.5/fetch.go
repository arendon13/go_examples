package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// go run fetch.go http://gopl.io
// go run fetch.go http://bad.gopl.io
// go run fetch.go http://bad.gopl.io
func main() {

	for _, userURL := range os.Args[1:] {

		url := checkForPrefix(userURL, "http://")

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status Code: %s\n", resp.Status)

		fmt.Print("\n\nResponse Body:\n")
		n, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		}

		fmt.Printf("\n\nResponse Body Count:\n%d", n)

	}

}

func checkForPrefix(url string, prefix string) string {

	if strings.HasPrefix(url, prefix) {
		return url
	}

	return prefix + url

}
