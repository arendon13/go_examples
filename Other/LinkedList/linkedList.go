package main

import (
	"fmt"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
}

func (ll *linkedList) insert(data int) {

	n := new(node)
	n.value = data
	n.next = nil

	if ll.head == nil {
		ll.head = n
	} else {
		curNode := ll.head
		for curNode.next != nil {
			curNode = curNode.next
		}
		curNode.next = n
	}

}

func (ll *linkedList) insertAtStart(data int) {

	n := new(node)
	n.value = data
	n.next = ll.head
	ll.head = n

}

func (ll *linkedList) insertAt(index, data int) {

	if index == 0 {
		ll.insertAtStart(data)
	} else {

		n := new(node)
		n.value = data

		curNode := ll.head
		for i := 0; i < index-1; i++ {
			curNode = curNode.next
		}

		n.next = curNode.next
		curNode.next = n

	}

}

func (ll *linkedList) deleteAt(index int) {

	if index == 0 {
		ll.head = ll.head.next
	} else {

		curNode := ll.head
		for i := 0; i < index-1; i++ {
			curNode = curNode.next
		}

		nodeToDelete := curNode.next
		curNode.next = nodeToDelete.next

		nodeToDelete = nil

	}

}

func (ll linkedList) String() string {

	var builder strings.Builder

	curNode := ll.head
	for {
		if curNode == nil {
			builder.WriteString(fmt.Sprintf("empty"))
			break
		} else if curNode.next == nil {
			builder.WriteString(fmt.Sprintf("%d", curNode.value))
			break
		} else {
			builder.WriteString(fmt.Sprintf("%d,", curNode.value))
			curNode = curNode.next
		}
	}

	return builder.String()

}

func main() {

	var ll linkedList
	ll.insert(12)
	ll.insert(87)
	ll.insert(9)
	ll.insertAtStart(24)
	ll.insert(44)
	ll.insertAt(2, 33)
	ll.insertAt(0, 3)
	fmt.Println(ll)

	ll.deleteAt(5)
	fmt.Println(ll)

}
