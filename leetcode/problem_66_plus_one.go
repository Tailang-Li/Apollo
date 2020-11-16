package leetcode

/*
Given a non-empty array of decimal digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.
*/

func plusOne(digits []int) []int {

	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 10 {

			digits[i] = 0
			if i == 0 {
				result := make([]int, 1)
				result[0] = 1
				result = append(result, digits...)
				return result
			}

			digits[i-1]++
		}
	}

	return digits
}
