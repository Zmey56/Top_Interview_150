package maxProfit

import "math"

//You are given an array prices where prices[i] is the price of a given stock on the ith day.
//
//You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
//
//Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func maxProfitVar0(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

//Time Limit Exceeded

func maxProfitVar1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	profit := 0
	minPrice := float64(prices[0])
	for i := 1; i < len(prices); i++ {
		minPrice = math.Min(minPrice, float64(prices[i]))
		profit = int(math.Max(float64(profit), float64(prices[i])-minPrice))
	}
	return profit
}

func maxProfitVar2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

func maxProfitVar3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	descending := true
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			descending = false
			break
		}
	}
	if descending {
		return 0
	}

	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
		} else {
			profit := price - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}
