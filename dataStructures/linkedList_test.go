package dataStructures

import "testing"

func TestAddNodeToEmptyList(t *testing.T) {
	data := 10
	var ll LinkedList
	if ll.Length() != 0 {
		t.Errorf("New Linked List length greater than 0")
	}

	if !ll.IsEmpty() {
		t.Errorf("New Linked List is not Empty")
	}

	ll.AddNode(data)

	if ll.IsEmpty() || ll.Length() != 1 {
		t.Errorf("Empty or Length greater than 1 after single node insertion")
	}

	if ll.root.data != data {
		t.Errorf("Data incorrectly stored")
	}
}
