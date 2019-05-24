package main

import (
	"flag"
	"fmt"
)

func main() {

	greetingPtr := flag.String("g", "Hello!", "a general greeting")
	wavePtr := flag.Bool("w", true, "decides whether or not a person will wave")

	flag.Parse()

	if *wavePtr {
		fmt.Printf("%s; You also wave!\n", *greetingPtr)
	} else {
		fmt.Println(*greetingPtr)
	}

}
