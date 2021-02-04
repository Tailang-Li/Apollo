package main

import "fmt"

/*
There are several cards arranged in a row, and each card has an associated number of points The points are given in the integer array cardPoints.

In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.

Your score is the sum of the points of the cards you have taken.

Given the integer array cardPoints and the integer k, return the maximum score you can obtain.
*/

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}

func maxScore(cardPoints []int, k int) int {

	var (
		windowSize      int = len(cardPoints) - k
		windowCnt       int
		windowTotal     int
		minScore        int
		totalCardPoints int
	)

	for i := range cardPoints {

		totalCardPoints += cardPoints[i]

		if windowCnt != windowSize {
			windowTotal += cardPoints[i]
			minScore = windowTotal
			windowCnt++
			continue
		}

		windowTotal = windowTotal + cardPoints[i] - cardPoints[i-windowSize]
		if windowTotal < minScore {
			minScore = windowTotal
		}
	}

	return totalCardPoints - minScore
}
