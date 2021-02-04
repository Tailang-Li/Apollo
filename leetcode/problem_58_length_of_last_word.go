package main

import "fmt"

/*
Given a string s consists of some words separated by spaces, return the length of the last word in the string. If the last word does not exist, return 0.

A word is a maximal substring consisting of non-space characters only.
*/

func main() {
	fmt.Println(lengthOfLastWord("hello world"))
	fmt.Println(lengthOfLastWord(" "))
	fmt.Println(lengthOfLastWord("a"))
}

func lengthOfLastWord(s string) int {

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] != ' ' {

			j := i
			for ; j >= 0; j-- {
				if s[j] == ' ' {
					return i - j
				}
			}

			return i + 1
		}
	}

	return 0
}
