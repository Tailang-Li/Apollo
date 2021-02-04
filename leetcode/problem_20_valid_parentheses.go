package main

import "fmt"

var (
	parenthesesPair map[rune]rune = map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
)

func main() {
	fmt.Println(isValid("{[()]}"))
}

func isValid(s string) bool {

	stack := make([]rune, 0)

	for _, p := range s {

		if _, ok := parenthesesPair[p]; !ok {
			stack = append(stack, p)
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != parenthesesPair[p] {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
