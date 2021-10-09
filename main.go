package main

import "fmt"

func printHello() {
	fmt.Println("Hello World")
}

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

func main() {
	var linkedList LinkedList
	linkedList.display()
	linkedList.addNode(0)
	linkedList.display()
	linkedList.addNode(1)
	linkedList.display()
	linkedList.addNodeAtBeginning(2)
	linkedList.display()
	fmt.Println()
	// var root Node
	// root.data = 0

	// var temp Node
	// temp.data = 1

	// root.next = &temp
	// var current *Node = &root
	// for current != nil {
	// 	fmt.Println(current.data)
	// 	current = current.next
	// }
}
