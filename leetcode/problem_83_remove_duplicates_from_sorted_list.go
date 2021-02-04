package main

import "fmt"

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}

	deleteDuplicates(&head)

	l := &head
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func deleteDuplicates(head *ListNode) *ListNode {

	l := head
	for l != nil && l.Next != nil {

		if l.Next.Val != l.Val {
			l = l.Next
			continue
		}

		tmp := l.Next
		l.Next = l.Next.Next
		tmp.Next = nil
	}

	return head
}
