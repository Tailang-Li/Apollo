package leetcode

/*
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it.
*/

func generate(numRows int) [][]int {

	result := make([][]int, 0, numRows)
	if numRows == 0 {
		return result
	}

	result = append(result, []int{1})

	for i := 1; i < numRows; i++ {

		line := make([]int, i+1)
		line[0] = 1
		line[i] = 1

		for j := 1; j < i; j++ {
			line[j] = result[i-1][j-1] + result[i-1][j]
		}

		result = append(result, line)
	}

	return result
}
