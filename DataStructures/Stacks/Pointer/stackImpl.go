package main

import "fmt"

type stack struct {
	list []int
}

func (s *stack) push(v int) {

	s.list = append(s.list, v)

}

func (s *stack) pop() int {

	size := len(s.list)

	s.list = s.list[:size-1]

	return s.list[size-2]

}

func (s *stack) peek() int {

	size := len(s.list)

	return s.list[size-1]

}

func (s stack) isEmpty() bool {

	return len(s.list) == 0

}

func main() {

	var stack stack

	fmt.Println(stack.isEmpty())

	stack.push(1)

	fmt.Println(stack.isEmpty())

	fmt.Println(stack.peek())

	stack.push(2)

	fmt.Println(stack.peek())

	fmt.Println(stack.list)

	_ = stack.pop()

	fmt.Println(stack.list)

}
