package main

import (
	"fmt"
	"sort"
)

/*
You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope.

One envelope can fit into another if and only if both the width and height of one envelope is greater than the width and height of the other envelope.

Return the maximum number of envelopes can you Russian doll (i.e., put one inside the other).

Note: You cannot rotate an envelope.
*/

func main() {
	fmt.Println(maxEnvelopes([][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}))
}

func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0]*envelopes[i][1] > envelopes[j][0]*envelopes[j][1]
	})

	var totalMax int
	envelopesCnt := make([]int, len(envelopes))
	for i := range envelopes {

		maxCnt := 1
		for j := i - 1; j >= 0; j-- {
			if canPutIn(envelopes[i], envelopes[j]) {
				maxCnt = max(maxCnt, envelopesCnt[j]+1)
				if envelopesCnt[j] == totalMax {
					break
				}
			}
		}

		totalMax = max(totalMax, maxCnt)
		envelopesCnt[i] = maxCnt
	}

	return totalMax
}

func canPutIn(a, b []int) bool {
	return a[0] < b[0] && a[1] < b[1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
