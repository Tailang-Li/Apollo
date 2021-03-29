package main

import "fmt"

/*
Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.
*/

func main() {
	l3 := ListNode{
		Val: 2,
	}
	l2 := ListNode{
		Val:  2,
		Next: &l3,
	}
	l1 := ListNode{
		Val:  1,
		Next: &l2,
	}

	rs := deleteDuplicates(&l1)

	for rs != nil {
		fmt.Println(rs.Val)
		rs = rs.Next
	}
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	dummy := &ListNode{}
	p := dummy
	for head != nil {

		switch {
		case head.Next == nil,
			head.Next != nil && head.Next.Val != head.Val:
			p.Next = head
		default:
			head = findNextDiff(head)
			continue
		}

		p = p.Next
		head = head.Next
		p.Next = nil
	}

	return dummy.Next
}

func findNextDiff(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	for head.Next != nil {
		if head.Val != head.Next.Val {
			return head.Next
		}
		head = head.Next
	}

	return nil
}
