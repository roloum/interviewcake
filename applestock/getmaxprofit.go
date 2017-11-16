package applestock

//GetMaxProfit ...
//Return the maximum profit one could make from a list of stock prices
func GetMaxProfit(stock []int) int {

	stockLen := len(stock)
	//Validate list size
	if stockLen < 2 {
		panic("Function requires list with at least 2 prices")
	}

	//Precompute the max profit with the first two values of the list
	if stock[0] < 0 {
		panic("Price can't be negative")
	}
	minPrice := stock[0]
	maxProfit := stock[1] - stock[0]

	//Iterate the list from index 1 and check if the currentProfit > maxProfit
	for i := 1; i < stockLen; i++ {

		if stock[i] < 0 {
			panic("Price can't be negative")
		}
		currentPrice := stock[i]
		currentProfit := currentPrice - minPrice

		//Update maxProfit
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}

		//Update minPrice
		if currentPrice < minPrice {
			minPrice = currentPrice
		}
	}

	return maxProfit
}
