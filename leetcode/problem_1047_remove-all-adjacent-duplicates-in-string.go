package main

import (
	"fmt"
	"strings"
)

/*
Given a string S of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.

We repeatedly make duplicate removals on S until we no longer can.

Return the final string after all such duplicate removals have been made.  It is guaranteed the answer is unique.
*/

func main() {
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("aaaaaaaaa"))
}

func removeDuplicates(S string) string {

	stack := make([]rune, 0)
	for _, r := range S {
		if len(stack) == 0 || stack[len(stack)-1] != r {
			stack = append(stack, r)
			continue
		}
		stack = stack[:len(stack)-1]
	}

	var b strings.Builder
	for _, r := range stack {
		b.WriteRune(r)
	}

	return b.String()
}
