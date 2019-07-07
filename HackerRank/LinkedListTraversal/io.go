package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	numScan, listScan := readUserInputForValues()

	num, numList := formatUserInput(numScan, listScan)

	fmt.Println(numList.Size())

	mElement, err := numList.mthToLast(num)
	if err != nil {
		fmt.Println("NIL")
		return
	}

	fmt.Println(mElement)
}

// Node represents an element in a Linked List
type Node struct {
	val  int
	next *Node
}

// LinkedList represents a object of Nodes
type LinkedList struct {
	head *Node
}

// Insert will insert a Node to the end of a Linked List
func (ll *LinkedList) Insert(num int) {

	// create a new node that will hold the value to be inserted
	n := new(Node)
	n.val = num
	n.next = nil

	// check if our head node exists
	// if head does not exist make the new node the head node
	// else traverse the list to the end and insert the new node
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

// Size will return the number of elements in a Linked List
func (ll *LinkedList) Size() int {

	if ll.head == nil {
		return 0
	}

	curNode := ll.head
	size := 1

	for curNode.next != nil {
		curNode = curNode.next
		size++
	}

	return size

}

func (ll *LinkedList) mthToLast(m int) (int, error) {
	// create two nodes that point to the head node
	n1 := ll.head
	n2 := ll.head

	// we want n2 to be m-1 nodes away from n1
	// move n2
	for i := 0; i < m-1; i++ {
		if n2 == nil {
			err := fmt.Errorf("error: invalid index")
			return -1, err
		}
		n2 = n2.next
	}

	// move n1 and n2 until n2 reaches the last node in the list
	for n2.next != nil {
		n1 = n1.next
		n2 = n2.next
	}

	return n1.val, nil
}

func (ll LinkedList) String() string {
	var stringer strings.Builder

	stringer.WriteString("Elements: ")

	curNode := ll.head
	for curNode.next != nil {
		stringer.WriteString(strconv.Itoa(curNode.val) + " ")
		curNode = curNode.next
	}
	stringer.WriteString(strconv.Itoa(curNode.val) + " ")

	return stringer.String()
}

func readUserInputForValues() (string, string) {

	scanner := bufio.NewScanner(os.Stdin)

	var numScan, listScan string
	if scanner.Scan() {
		numScan = scanner.Text()
	}
	if scanner.Err() != nil {
		panic(1)
	}

	if scanner.Scan() {
		listScan = scanner.Text()
	}
	if scanner.Err() != nil {
		panic(1)
	}

	return numScan, listScan

}

func formatUserInput(num, list string) (int, LinkedList) {
	m, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal(err)
	}

	listSlice := strings.Split(list, " ")

	var newList LinkedList
	for _, n := range listSlice {
		e, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}

		newList.Insert(e)
	}

	return m, newList
}
