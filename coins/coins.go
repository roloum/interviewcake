package coins

//GetNumberOfWaysToMakeAmountWithCoins ...
func GetNumberOfWaysToMakeAmountWithCoins(coins []int, amount int) int {

	//Array stores how many ways we can make an amount with coins, going from 1
	//to the actual amount. Array key is the amount, value is the number of coins
	var waysOfMakingNCents = make([]int, amount+1)
	//Position 0 represents how many ways of making amount X with a coin of the
	//same amount,
	waysOfMakingNCents[0] = 1

	for _, coin := range coins {

		//Figure out how many ways of making all different amounts with each of the
		//coins, starting with the amount of the coin.
		for tempAmount := coin; tempAmount <= amount; tempAmount++ {
			remainder := tempAmount - coin
			waysOfMakingNCents[tempAmount] += waysOfMakingNCents[remainder]
		}
	}
	return waysOfMakingNCents[amount]

}
