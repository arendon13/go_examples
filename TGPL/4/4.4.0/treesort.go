package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	someList := []int{3, 6, 1, 34, 2, 7, 4}
	fmt.Println(someList)
	sort(someList)
	fmt.Println(someList)
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// (Depth First search: In Order Traversal: Left-Root-Right)
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
