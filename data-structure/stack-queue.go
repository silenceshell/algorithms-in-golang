package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stack struct {
	head *Node
}

type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

func (stack *Stack) Push(value interface{}) {
	node := &Node{next: stack.head, value: value}
	stack.head = node
}

func (stack *Stack) Pop() (value interface{}) {
	node := stack.head
	stack.head = stack.head.next
	return node.value
}

func (stack *Stack) Display() {
	node := stack.head
	fmt.Print("stack top:")
	for node != nil {
		fmt.Printf(" => %d", node.value)
		node = node.next
	}
	fmt.Println(" :stack bottom")
}

func (stack *Stack) IsEmpty() bool {
	if stack.head == nil {
		return true
	} else {
		return false
	}
}

type Queue struct {
	head *Node
	tail *Node
}

func (queue *Queue) Push(value interface{}) {
	var node *Node = &Node{next: queue.head, prev: nil, value: value}
	if queue.head != nil {
		queue.head.prev = node
	}
	queue.head = node

	tmp := queue.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	queue.tail = tmp
}

func (queue *Queue) Pop() (value interface{}) {
	node := queue.tail
	queue.tail = queue.tail.prev
	if queue.tail != nil {
		queue.tail.next = nil
	}

	return node.value
}

func (queue *Queue) IsEmpty() bool {
	if queue.tail == nil {
		return true
	} else {
		return false
	}
}

func (queue *Queue) Display() {
	node := queue.tail
	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.prev
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var stack Stack = Stack{}

	for i := 0; i < 10; i++ {
		stack.Push(rand.Intn(999) - rand.Intn(999))
	}
	stack.Display()
	for !stack.IsEmpty() {
		fmt.Printf("%d ", stack.Pop())
	}
	fmt.Println()

	var queue Queue = Queue{}

	for i := 0; i < 10; i++ {
		queue.Push(rand.Intn(999) - rand.Intn(999))
	}
	queue.Display()
	fmt.Println()

	for !queue.IsEmpty() {
		fmt.Printf("%d ", queue.Pop())
	}
	fmt.Println()
}
