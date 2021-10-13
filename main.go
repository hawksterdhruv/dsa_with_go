package main

import (
	"fmt"

	"github.com/hawksterdhruv/dsa_with_go/dataStructures"
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
}
