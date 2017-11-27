package main

import "fmt"

// https://zh.wikipedia.org/wiki/%E9%93%BE%E8%A1%A8#/media/File:Doubly-linked-list.svg
type Node struct {
	next *Node
	prev *Node
	key interface{}
}

type List struct{
	head *Node
	tail *Node
}

func (list * List) insert(key interface{}) {
	node := &Node{next:list.head, prev:nil, key:key}
	if list.head != nil {
		list.head.prev = node
	}
	list.head = node

	tmp := list.head
	// todo: is this needed check every time?
	for tmp.next != nil {
		tmp = tmp.next
	}
	list.tail = tmp
}

func (list *List) display() {
	node := list.head
	for node != nil {
		fmt.Printf(" %v ", node.key)
		node = node.next
	}
	fmt.Println()
}

func (list *List) reverse() {
	node := list.head

	for node != nil {
		node.next, node.prev = node.prev, node.next
		node = node.prev
	}

	list.head, list.tail = list.tail, list.head
}

func main() {
	list := List{}
	list.insert(12)
	list.insert(44)
	list.insert(11)
	list.insert(-11)
	list.insert(552)
	list.insert(34)
	list.insert(-911)
	list.insert(-11)

	list.display()

	list.reverse()
	list.display()
}
