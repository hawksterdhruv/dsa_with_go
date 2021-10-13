package dataStructures

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBST_Insert(t *testing.T) {
	bst := BST{}
	if bst.root != nil {
		t.Errorf("New BST isn't empty")
	}
	bst.Insert(32)
	bst.Insert(16)
	bst.Insert(48)
	bst.Insert(1)
	bst.Insert(17)

	if bst.root.data != 32 {
		t.Errorf("New node mismatch")
	}
	if bst.root.left.data != 16 {
		t.Errorf("Left Insertion error")
	}
	if bst.root.right.data != 48 {
		t.Errorf("Right insertion error")
	}
	if bst.root.left.left.data != 1 {
		t.Errorf("Left Left Insertion error")
	}
	if bst.root.left.right.data != 17 {
		t.Errorf("Left Right Insertion error")
	}

}

func TestBST_InOrderTraversal(t *testing.T) {
	bst := BST{}
	if bst.root != nil {
		t.Errorf("New BST isn't empty")
	}

	_, err := bst.InOrderTraversal()

	fmt.Printf("%v : %T", err, err)
	if err == nil {
		t.Errorf("No error on empty BST")
	}

	bst.Insert(32)
	bst.Insert(16)
	bst.Insert(48)
	bst.Insert(1)
	bst.Insert(17)

	actual, err := bst.InOrderTraversal()
	if err != nil {
		t.Errorf("InOrder traversal failed, %s", err)
	}
	expected := []int{1, 16, 17, 32, 48}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Inorder tranversal didn't match")
	}

}
