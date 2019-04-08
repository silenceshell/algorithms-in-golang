package main

import "fmt"

func main() {
	var l1 *ListNode = new(ListNode).Set(2)
	l1.Append(new(ListNode).Set(4))
	l1.Append(new(ListNode).Set(3))
	l1.Print()

	var l2 *ListNode = new(ListNode).Set(5)
	l2.Append(new(ListNode).Set(6))
	l2.Append(new(ListNode).Set(4))
	l2.Print()

	var l3 *ListNode = addTwoNumbers(l1, l2)
	l3.Print()
}

type ListNode struct {
     Val int
     Next *ListNode
}

func (node *ListNode) Append(l *ListNode) {
	for node.Next != nil {
		node = node.Next
	}
	node.Next = l
}

func (node *ListNode) Set(val int) *ListNode {
	node.Val = val
	return node
}

func (node *ListNode) Print() {
	var l *ListNode = node
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Println("========")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var  last *ListNode
	var v int = 0
	var x int = 0
	var up int = 0
	var b1 bool = false
	var b2 bool = false
	for {
		v = up
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		} else {
			b1 = true
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		} else {
			b2 = true
		}

		if b1 && b2 && up==0 {
			break
		}

		n := new(ListNode)
		n.Val = v%10
		up = v/10

		if x == 0 {
			result = n
			last = n
			x = 1
		} else {
			last.Next = n
			last = n
		}

	}

	return result
}