package main

import "fmt"

/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

func main() {

	head := ListNode{
		Val: 3,
	}

	l1 := ListNode{
		Val: 2,
	}

	l2 := ListNode{
		Val: 0,
	}

	l3 := ListNode{
		Val: -4,
	}

	head.Next = &l1
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l1

	fmt.Println(hasCycle(&head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil {
		return false
	}

	fast, slow := head, head
	for {

		// move fast pointer
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next

		// move slow pointer
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
}
