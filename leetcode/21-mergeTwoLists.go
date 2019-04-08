package main
import (
    "fmt"
)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func (l *ListNode) append(node *ListNode) {
    for l.Next != nil {
        l = l.Next
    }
    l.Next = node
}
func (l *ListNode) print() {
    n := l
    for n != nil {
        fmt.Println(n.Val)
        n = n.Next
    }
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var head ListNode
    index := &head
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            index.Next = l1
            l1 = l1.Next
        } else {
            index.Next = l2
            l2 = l2.Next
        }
        index = index.Next
    }
    if l1 != nil {
        index.Next = l1
    }
    if l2 != nil {
        index.Next = l2
    }
    return head.Next
}

func main() {
    l1 := &ListNode{1, nil}
    l1.append(&ListNode{2, nil})
    l1.append(&ListNode{3, nil})
    l1.append(&ListNode{5, nil})
    l1.append(&ListNode{10, nil})
    l1.print()
    fmt.Println("---------")

    l2 := &ListNode{2, nil}
    l2.append(&ListNode{4, nil})
    l2.append(&ListNode{8, nil})
    l2.append(&ListNode{9, nil})
    l2.append(&ListNode{10, nil})
    l2.append(&ListNode{20, nil})
    l2.print()
    fmt.Println("---------")

    l := mergeTwoLists(l1, l2)
    l.print()
    fmt.Println("---------")

    l = mergeTwoLists(&ListNode{}, &ListNode{})
    l.print()
    fmt.Println("---------")

    l = mergeTwoLists(nil, nil)
    l.print()
    fmt.Println("---------")

}
