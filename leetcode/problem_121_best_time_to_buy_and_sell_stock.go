package leetcode

/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.
*/

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	result, min := 0, prices[0]

	for i := 1; i < len(prices); i++ {

		if prices[i] < min {
			min = prices[i]
		}

		if prices[i]-min > result {
			result = prices[i] - min
		}
	}

	return result
}
