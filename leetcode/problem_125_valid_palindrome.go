package main

import "fmt"

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.
*/

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("0P"))
}

func isPalindrome(s string) bool {

	if s == "" {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {

		// left move until meet a alphanumeric char
		for {

			if left == right {
				return true
			}

			if isAlphanumeric(s[left]) {
				break
			}

			left++
		}

		// right move until meet a alphanumeric char
		for {

			if left == right {
				return true
			}

			if isAlphanumeric(s[right]) {
				break
			}

			right--
		}

		if !equal(s[left], s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func equal(a, b byte) bool {

	if a == b {
		return true
	}

	if a >= 'a' && a <= 'z' {
		return int(a-b) == 32
	}

	if b >= 'a' && b <= 'z' {
		return int(b-a) == 32
	}

	return false
}

func isAlphanumeric(b byte) bool {
	return b >= '0' && b <= '9' ||
		b >= 'a' && b <= 'z' ||
		b >= 'A' && b <= 'Z'
}
