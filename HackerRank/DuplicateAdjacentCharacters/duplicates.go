package main

import (
	"fmt"
)

func main() {

	s1 := "abccde"
	s2 := "abccbe"
	s3 := "absbe"
	s4 := "abccba"
	s5 := "aabcde"

	fmt.Println(removeDuplicates(s1))
	fmt.Println(removeDuplicates(s2))
	fmt.Println(removeDuplicates(s3))
	fmt.Println(removeDuplicates(s4))
	fmt.Println(removeDuplicates(s5))

}

type stack []byte

func (s stack) push(v byte) stack {

	return append(s, v)

}

func (s stack) pop() (stack, byte) {

	size := len(s)

	return s[:size-1], s[size-1]

}

func (s stack) peek() byte {

	size := len(s)

	return s[size-1]

}

func (s stack) isEmpty() bool {

	return len(s) == 0

}

func removeDuplicates(s string) string {

	var stack stack

	stringInBytes := []byte(s)

	previous := stringInBytes[0]

	stack = stack.push(previous)

	// var err error

	for i := 1; i < len(stringInBytes); i++ {

		current := stringInBytes[i]

		if current != previous {
			// add current character to stack
			stack = stack.push(current)
			previous = current
		} else {
			stack, _ = stack.pop()
			if !stack.isEmpty() {
				previous = stack.peek()
			}
		}

	}

	return string(stack)

}
