package main

import "fmt"

/*
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
*/

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	s1InfoMap := make(map[byte]int)
	s2InfoMap := make(map[byte]int)
	for i := range s1 {
		s1InfoMap[s1[i]] = s1InfoMap[s1[i]] + 1
		s2InfoMap[s2[i]] = s2InfoMap[s2[i]] + 1
	}

	if checkEqual(s1InfoMap, s2InfoMap) {
		return true
	}

	for i := len(s1); i < len(s2); i++ {

		left := i - len(s1)

		s2InfoMap[s2[left]] = s2InfoMap[s2[left]] - 1
		if s2InfoMap[s2[left]] == 0 {
			delete(s2InfoMap, s2[left])
		}

		s2InfoMap[s2[i]] = s2InfoMap[s2[i]] + 1

		if checkEqual(s1InfoMap, s2InfoMap) {
			return true
		}
	}

	return false
}

func checkEqual(map1, map2 map[byte]int) bool {

	if len(map1) != len(map2) {
		return false
	}

	for k, v := range map1 {
		if map2[k] != v {
			return false
		}
	}

	return true
}
