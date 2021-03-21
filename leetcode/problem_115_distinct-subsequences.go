package main

import "fmt"

/*
Given two strings s and t, return the number of distinct subsequences of s which equals t.

A string's subsequence is a new string formed from the original string by deleting some (can be none) of the characters without disturbing the remaining characters' relative positions. (i.e., "ACE" is a subsequence of "ABCDE" while "AEC" is not).

It is guaranteed the answer fits on a 32-bit signed integer.
*/

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
}

func numDistinct(s string, t string) int {

	m, n := len(s), len(t)
	// s[i:] t[j:]
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {

			if m-i < n-j {
				break
			}

			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

func calcDistinct(s string, t string) int {

	if len(t) > len(s) || len(t) == 0 || len(s) == 0 {
		return 0
	}

	var count int
	if len(t) == 1 {
		for i := range s {
			if s[i] == t[0] {
				count++
			}
		}
		return count
	}

	for i := 0; i <= len(s)-len(t); i++ {
		if s[i] == t[0] {
			count += numDistinct(s[i+1:], t[1:])
		}
	}

	return count
}
