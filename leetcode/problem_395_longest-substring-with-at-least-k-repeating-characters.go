package main

import (
	"fmt"
	"strings"
)

/*
Given a string s and an integer k, return the length of the longest substring of s such that the frequency of each character in this substring is greater than or equal to k.
*/

func main() {
	fmt.Println(longestSubstring("ababbc", 2))
	fmt.Println(longestSubstring("aaabb", 3))
}

func longestSubstring(s string, k int) int {

	letterCnts := make([]int, 26)
	for i := range s {
		letterCnts[s[i]-'a']++
	}

	var split byte
	for i, letterCnt := range letterCnts {
		if letterCnt > 0 && letterCnt < k {
			split = byte(i) + 'a'
		}
	}

	if split == 0 {
		return len(s)
	}

	var result int
	for _, subString := range strings.Split(s, string(split)) {
		result = max(result, longestSubstring(subString, k))
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
