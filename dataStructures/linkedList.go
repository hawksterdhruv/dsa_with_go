package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	root *Node
}

func (ll *LinkedList) addNode(data int) {
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

func (ll *LinkedList) addNodeAtBeginning(data int) {
	node := Node{data: data}
	if ll.root == nil {
		ll.root = &node
	} else {
		node.next = ll.root
		ll.root = &node
	}
}

func (ll LinkedList) display() {
	current := ll.root
	fmt.Println()
	for current != nil {
		fmt.Printf("%v ->", current.data)
		current = current.next
	}
}
