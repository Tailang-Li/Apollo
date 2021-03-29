package main

import "fmt"

/*
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.
*/

func main() {
	fmt.Println("vim-go")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	dummy := head
	for head != nil {

		if head.Next != nil && head.Next.Val == head.Val {
			tmp := head.Next
			head.Next = head.Next.Next
			tmp.Next = nil
			continue
		}

		head = head.Next
	}

	return dummy
}
