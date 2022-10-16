package main

func buySellStock_Brute(prices []int) int {

	var max int
	for i, p := range prices {
		for _, t := range prices[i+1:] {
			if t-p > max {
				max = t - p
			}
		}
	}
	return max
}

/*
If we can keep track of the minimum stock price and maximum profit,
we should be able to solve the problem in single pass.

https://medium.com/algorithms-and-leetcode/best-time-to-buy-sell-stocks-on-leetcode-the-ultimate-guide-ce420259b323
*/

func buySellStock_Optimized(prices []int) int {

	var max int
	min_price := prices[0]

	for _, p := range prices {
		temp := p - min_price
		if temp > max {
			max = temp
		}

		if p < min_price {
			min_price = p
		}

	}

	return max
}
