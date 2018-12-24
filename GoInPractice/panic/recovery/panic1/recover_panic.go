package main

import (
	"errors"
	"fmt"
)

// technique 20
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	yikes()
}

func yikes() {
	panic(errors.New("something bad happened"))
}

// the following would also work
// func callRecover() {
// 	if err := recover(); err != nil {
// 		fmt.Printf("Trapped panic: %s (%T)\n", err, err)
// 	}
// }
