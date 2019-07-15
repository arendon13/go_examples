package main

import "fmt"

type stack []int

func (s stack) push(v int) stack {

	return append(s, v)

}

func (s stack) pop() (stack, int) {

	size := len(s)

	return s[:size-1], s[size-1]

}

func (s stack) peek() int {

	size := len(s)

	return s[size-1]

}

func (s stack) isEmpty() bool {

	return len(s) == 0

}

func main() {

	var stack stack

	fmt.Println(stack.isEmpty())

	stack = stack.push(1)

	fmt.Println(stack.isEmpty())

	fmt.Println(stack.peek())

	stack = stack.push(2)

	fmt.Println(stack.peek())

	fmt.Println(stack)

}
