package main

import (
	"fmt"
	"os"
	"time"
)

// technique 13
func main() {
	done := time.After(30 * time.Second) // returns a "read only channel" <-chan
	echo := make(chan []byte)
	go readStdin(echo)
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("\nTimed out")
			os.Exit(0)
		}
	}
}

func readStdin(out chan<- []byte) { // channel in this function is "write only channel" chan<-
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}
