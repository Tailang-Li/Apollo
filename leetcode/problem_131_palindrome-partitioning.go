package main

import "fmt"

/*
Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.

A palindrome string is a string that reads the same backward as forward.
*/

func main() {
	fmt.Println(partition("aab"))
}

var (
	checkArr    [][]bool
	result      [][]string
	tempStrings []string
)

func partition(s string) [][]string {

	checkArr = make([][]bool, len(s))
	result = make([][]string, 0)
	tempStrings = make([]string, 0)

	for i := range s {
		checkArr[i] = make([]bool, len(s))
		for j := range s {
			checkArr[i][j] = true
		}
	}

	for i := 1; i < len(s); i++ {
		p, q := 0, i
		for q < len(s) {
			checkArr[p][q] = s[p] == s[q] && checkArr[p+1][q-1]
			p++
			q++
		}
	}

	dfs(0, s)

	return result
}

func dfs(i int, s string) {

	if i == len(s) {
		result = append(result, append([]string(nil), tempStrings...))
		return
	}

	for j := i; j < len(s); j++ {
		if checkArr[i][j] {
			tempStrings = append(tempStrings, s[i:j+1])
			dfs(j+1, s)
			tempStrings = tempStrings[:len(tempStrings)-1]
		}
	}
}
