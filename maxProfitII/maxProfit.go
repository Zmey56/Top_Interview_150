package maxProfitII

import "math"

// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
//
//On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time.
//However, you can buy it then immediately sell it on the same day.
//
//Find and return the maximum profit you can achieve.

func maxProfitVar0(prices []int) int {
	profit := 0
	minPrice := float64(prices[0])
	for i := 1; i < len(prices); i++ {
		minPrice = math.Min(minPrice, float64(prices[i]))
		diff := float64(prices[i]) - minPrice
		if diff > 0 {
			profit += int(diff)
			minPrice = float64(prices[i])
		}
	}
	return profit
}

func maxProfitVar1(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit := prices[i] - minPrice
			maxProfit += profit
			minPrice = prices[i]
		}
	}
	return maxProfit
}

func maxProfitVar2(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func calculateProfit(prices []int, start, end int, profitChan chan int) {
	profit := 0
	for i := start; i < end; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	profitChan <- profit
}

func maxProfitVar3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	profitChan := make(chan int)
	chunkSize := len(prices) / 4
	remainder := len(prices) % 4

	start := 1
	for i := 0; i < 4; i++ {
		end := start + chunkSize
		if remainder > 0 {
			end++
			remainder--
		}
		if end > len(prices) {
			end = len(prices)
		}
		go calculateProfit(prices, start, end, profitChan)
		start = end

	}

	totalProfit := 0
	for i := 0; i < 4; i++ {
		totalProfit += <-profitChan
	}

	return totalProfit

}
