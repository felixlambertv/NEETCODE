package neetcode

import "testing"

func maxProfit(prices []int) int {
	minPrice := prices[0]
	max := 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > max {
				max = profit
			}
		}
	}
	return max
}

func TestMaxProfit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	output := 5

	res := maxProfit(input)
	if res != output {
		t.Fail()
	}
}
