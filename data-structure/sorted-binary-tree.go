package main

import (
	"fmt"
	"github.com/silenceshell/algorithms-in-golang/utils"
)

type SortedBinaryNode struct {
	left *SortedBinaryNode
	right *SortedBinaryNode
	data int
}

type SortedBinaryTree struct {
	root *SortedBinaryNode
}

func (tree *SortedBinaryTree) insert(num int) *SortedBinaryTree {
	if tree.root == nil {
		tree.root = &SortedBinaryNode{left:nil, right:nil, data:num}
	} else {
		tree.root.insert(num)
	}
	return tree
}

func (node *SortedBinaryNode) insert(num int) error {
	if node == nil {
		return fmt.Errorf("node is nil, insert %d failed.", num)
	}
	if num <= node.data {
		if node.left == nil {
			node.left = &SortedBinaryNode{left:nil, right:nil, data:num}
		} else {
			node.left.insert(num)
		}
	} else {
		if node.right == nil {
			node.right = &SortedBinaryNode{left:nil, right:nil, data:num}
		} else {
			node.right.insert(num)
		}
	}
	return nil
}

func (tree *SortedBinaryTree) print() {
	tree.root.print()
}

func (node *SortedBinaryNode) print() {
	if node.left != nil {
		node.left.print()
	}
	fmt.Printf("%d ", node.data)
	if node.right != nil {
		node.right.print()
	}
}

func main() {
	tree := &SortedBinaryTree{}
	data := utils.GenerateSlice(13)
	fmt.Println(data)
	for _, v := range data {
		tree.insert(v)
	}
	tree.print()
}
