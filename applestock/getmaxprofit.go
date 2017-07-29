package applestock

import "fmt"

//GetMaxProfit ...
//Return the maximum profit one could make from a list of stock prices
func GetMaxProfit(stock []int) int {
	if len(stock) < 2 {
		panic("Getting a profit requires at least two prices")
	}
	minPrice := stock[0]
	maxProfit := stock[1] - stock[0]
	minLoss := maxProfit

	for _, price := range stock {
		if price <= 0 {
			panic(fmt.Sprintf("Invalid price %v", price))
		}
		newProfit := price - minPrice
		if newProfit > maxProfit {
			maxProfit = newProfit
		} else if newProfit < 0 {
			if newProfit > minLoss {
				minLoss = newProfit
			}
			minPrice = price
		}

	}

	if maxProfit > 0 {
		return maxProfit
	}
	return minLoss

}
