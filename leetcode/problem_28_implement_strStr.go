package main

import "fmt"

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := range haystack {
		if haystack[i] == needle[0] {

			found := true
			for j := range needle {

				if i+j == len(haystack) {
					return -1
				}

				if haystack[i+j] != needle[j] {
					found = false
					break
				}
			}

			if found {
				return i
			}
		}
	}

	return -1
}
