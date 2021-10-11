package main

import (
	"fmt"
	"github.com/hawksterdhruv/dataStructures"
)

func main() {
	var linkedList dataStructures.LinkedList
	linkedList.Display()
	linkedList.AddNode(0)
	linkedList.Display()
	linkedList.AddNode(1)
	linkedList.Display()
	linkedList.AddNodeAtBeginning(2)
	linkedList.Display()
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
