package dataStructures

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	root *Node
}

func (ll *LinkedList) AddNode(data int) {
	node := Node{data: data}
	if ll.root == nil {
		ll.root = &node
	} else {
		current := ll.root
		for current.next != nil {
			current = current.next
		}
		current.next = &node
	}
}

func (ll *LinkedList) AddNodeAtBeginning(data int) {
	node := Node{data: data}
	if ll.root == nil {
		ll.root = &node
	} else {
		node.next = ll.root
		ll.root = &node
	}
}

func (ll LinkedList) Display() {
	current := ll.root
	fmt.Println()
	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
}

func (ll LinkedList) IsEmpty() bool {
	return ll.root == nil
}

func (ll LinkedList) Length() int {
	current := ll.root
	count := 0
	for current != nil {
		current = current.next
		count++
	}
	return count
}
