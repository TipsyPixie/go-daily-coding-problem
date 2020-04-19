package problem047

func MaximizeProfit(prices []int) int {
	maxProfit, _ := getMaxProfitWithMaxPrice(prices)
	return maxProfit
}

func getMaxProfitWithMaxPrice(prices []int) (int, int) {
	switch len(prices) {
	case 0:
		return 0, 0
	case 1:
		return 0, prices[0]
	default:
		futureMaxProfit, futureMaxPrice := getMaxProfitWithMaxPrice(prices[1:])
		if currentPrice := prices[0]; currentPrice > futureMaxPrice {
			return futureMaxProfit, currentPrice
		} else if futureMaxPrice-currentPrice > futureMaxProfit {
			return futureMaxPrice - currentPrice, futureMaxPrice
		} else {
			return futureMaxProfit, futureMaxPrice
		}
	}
}
