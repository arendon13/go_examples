package main

import (
	"log"
	"net"
	"time"
)

// technique 24 - avoiding back pressure when logging
// run with 'nc -lk 1902'
func main() {

	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example", f)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
