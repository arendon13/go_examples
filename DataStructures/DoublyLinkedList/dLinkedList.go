package main

import (
	"fmt"
	"strings"
)

func main() {

	var dll doublyLinkedList

	dll.insert(2)
	dll.insert(4)
	dll.insertAtStart(1)
	dll.insertAt(2, 3)
	fmt.Println(dll)

	dll.deleteAt(2)
	fmt.Println(dll)

	dll.deleteAt(2)
	fmt.Println(dll)

}

type node struct {
	value int
	next  *node
	prev  *node
}

type doublyLinkedList struct {
	head *node
}

func (dll *doublyLinkedList) insert(data int) {

	n := new(node)
	n.value = data
	n.next = nil
	n.prev = nil

	if dll.head == nil {
		dll.head = n
	} else {
		curNode := dll.head
		for curNode.next != nil {
			curNode = curNode.next
		}
		curNode.next = n
		n.prev = curNode
	}

}

func (dll *doublyLinkedList) insertAtStart(data int) {

	n := new(node)
	n.value = data
	n.next = dll.head
	n.prev = nil

	dll.head.prev = n
	dll.head = n

}

func (dll *doublyLinkedList) insertAt(index, data int) {

	if index == 0 {
		dll.insertAtStart(data)
	} else {

		n := new(node)
		n.value = data

		curNode := dll.head
		for i := 0; i < index-1; i++ {
			curNode = curNode.next
		}

		n.next = curNode.next
		n.prev = curNode
		n.next.prev = n
		curNode.next = n

	}

}

func (dll *doublyLinkedList) deleteAt(index int) {

	if index == 0 {
		dll.head = dll.head.next
		dll.head.prev = nil
	} else {

		curNode := dll.head
		for i := 0; i < index-1; i++ {
			curNode = curNode.next
		}

		nodeToDelete := curNode.next
		if nodeToDelete.next != nil {
			curNode.next = nodeToDelete.next
			nodeToDelete.next.prev = curNode
		} else {
			curNode.next = nil
		}

	}

}

func (dll doublyLinkedList) String() string {

	var builder strings.Builder

	curNode := dll.head
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
