package main

import (
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)
}
