package main

import (
	"fmt"
	"strings"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go sentenceReader(chan1)
	go splitSentences(chan1, chan2)
	printWords(chan2)
}

func sentenceReader(downstream chan string) {
	for _, v := range []string{"hello world", "LA is black and gold", "a bad apple", "an eye for an eye"} {
		downstream <- v
	}
	close(downstream)
}

func splitSentences(upstream, downstream chan string) {
	for v := range upstream {
		words := strings.Split(v, " ")
		for _, word := range words {
			downstream <- word
		}
	}
	close(downstream)
}

func printWords(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
