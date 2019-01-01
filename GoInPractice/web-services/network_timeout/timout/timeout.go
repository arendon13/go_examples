package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// technique 49
func main() {
	_, err := http.Get("http://example.com/test.zip")
	if err != nil && hasTimedOut(err) {
		fmt.Println("A timeout error occured")
		return
	}
}

func hasTimedOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	case net.Error:
		if err.Timeout() {
			return true
		}
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	}

	errTxt := "use of closed network"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}
