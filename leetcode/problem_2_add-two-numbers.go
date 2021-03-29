package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("vim-go")
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	p := head
	var carry bool

	for l1 != nil || l2 != nil {

		var sum int

		switch {
		case l1 != nil && l2 != nil:
			sum = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
		case l1 != nil:
			sum = l1.Val
			l1 = l1.Next
		case l2 != nil:
			sum = l2.Val
			l2 = l2.Next
		}

		if carry {
			sum++
		}
		carry = sum >= 10
		p.Next = &ListNode{
			Val: sum % 10,
		}
		p = p.Next
	}

	if carry {
		p.Next = &ListNode{
			Val: 1,
		}
	}

	return head.Next
}
