package main

import (
	"errors"
	"time"

	"github.com/Masterminds/cookoo/safely"
)

// technique 21
func message() {
	println("Inside goroutine")
	panic(errors.New("oops"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(10000)
}
