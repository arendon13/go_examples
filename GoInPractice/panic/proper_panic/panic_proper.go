package main

import (
	"errors"
	"fmt"
)

/// technique 19
func main() {
	fmt.Println("Test")
	panic(errors.New("something bad happened"))
}
