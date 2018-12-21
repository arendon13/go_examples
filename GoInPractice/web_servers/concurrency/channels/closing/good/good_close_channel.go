package main

import (
	"fmt"
	"time"
)

// technique 14
func main() {
	msg := make(chan string)
	done := make(chan bool)
	until := time.After(5 * time.Second)

	go send(msg, done)

	// this is the "receiver" go routine
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true // if the receiver hits a stopping condition, it has to let the sender know
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) { // ch is a receiving/write-only channel, while done is a sending/read-only channel
	// this is the "sender" go routine
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Microsecond)
		}
	}
}
