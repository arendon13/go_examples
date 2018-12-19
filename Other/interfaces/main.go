package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// If receiver value is not being used, it can be removed.
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a spanish greeting
	return "Hola!"
}
