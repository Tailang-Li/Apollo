package main

import "fmt"

/*
Given a string s, partition s such that every substring of the partition is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.
*/

func main() {
	fmt.Println(minCut("aab"))
	fmt.Println(minCut("ab"))
	fmt.Println(minCut("a"))
	fmt.Println(minCut("leet"))
}

func minCut(s string) int {

	n := len(s)
	checkArr := make([][]bool, n)
	for i := range checkArr {
		checkArr[i] = make([]bool, n)
		for j := range checkArr[i] {
			checkArr[i][j] = true
		}
	}

	for i := 0; i < n-1; i++ {
		p, q := 0, i+1
		for q < n {
			checkArr[p][q] = s[p] == s[q] && checkArr[p+1][q-1]
			p++
			q++
		}
	}

	dp := make([]int, n)
	for i := range dp {
		if checkArr[0][i] {
			continue
		}
		dp[i] = n
		for j := 0; j < i; j++ {
			if checkArr[j+1][i] && dp[j]+1 < dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	return dp[n-1]
}
