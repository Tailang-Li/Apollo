package main

import "fmt"

/*
Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the first two lists.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	l := mergeTwoLists(l1, l2)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	switch {
	case l1 == nil:
		return l2
	case l2 == nil:
		return l1
	}

	head := &ListNode{}
	next := head
	p := l1
	q := l2
	for p != nil || q != nil {

		if p == nil || (q != nil && p.Val > q.Val) {
			next.Next = q
			next = next.Next
			q = q.Next
			continue
		}

		next.Next = p
		next = next.Next
		p = p.Next
	}

	return head.Next
}
