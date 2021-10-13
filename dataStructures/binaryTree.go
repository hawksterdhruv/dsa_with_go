package dataStructures

import "fmt"

type BstNode struct {
	data  int
	left  *BstNode
	right *BstNode
}

type BST struct {
	root *BstNode
}

func (a BST) Search() {

}

func (a BST) Mininum() {

}

func (a BST) Maximum() {

}

func (a BST) Predecessor() {

}

func (a BST) Successor() {

}

func (a *BST) Insert(data int) {
	node := BstNode{data: data}
	if a.root == nil {
		a.root = &node
	} else {
		current := a.root
		for {
			if data < current.data {
				if current.left != nil {
					current = current.left
				} else {
					current.left = &node
					break
				}
			} else {
				if current.right != nil {
					current = current.right
				} else {
					current.right = &node
					break
				}
			}
		}
	}
}

func (a BST) Delete() {

}

func (a BST) InOrderTraversal() ([]int, error) {
	if a.root != nil {
		val := visit(a.root)
		return val, nil
	}
	return nil, fmt.Errorf("BST empty")
}

func visit(node *BstNode) []int {
	var ar []int
	var left []int
	var right []int
	var val int
	if node == nil {
		return ar
	} else {
		left = visit(node.left)
		val = node.data
		right = visit(node.right)
	}
	ar = append(ar, left...)
	ar = append(ar, val)
	ar = append(ar, right...)
	return ar
}
