package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// technique 16
func main() {
	args := os.Args[1:]

	if result, err := Concat(args...); err != nil {
		// handles error
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}
	// because of the pattern used in the Concat function, the following
	// code could be run safely if the user does not care about the error
	result, _ := Concat(args...)
	fmt.Printf("Concatenated string: '%s'\n", result)
}

// Concat concatenates a bunch of strings, separated by spaces.Concat
// It returns an empty string and an error if no strings were passed in.
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied") // it is good practice to always return a useable value
	}

	return strings.Join(parts, " "), nil
}
