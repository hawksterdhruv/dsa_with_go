package main

import (
	"fmt"
)

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
