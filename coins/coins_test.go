package coins

import "testing"

func TestGetNumberOfWaysToMakeAmountWithCoins(t *testing.T) {
	cases := []struct {
		coins          []int
		amount, result int
	}{
		{[]int{1, 2, 3}, 4, 4},
		{[]int{1, 3, 5}, 5, 3},
	}

	for _, c := range cases {
		result := GetNumberOfWaysToMakeAmountWithCoins(c.coins, c.amount)
		if result != c.result {
			t.Errorf("Error! Coins: %v, amount: %v, expected: %v, received: %v",
				c.coins, c.amount, c.result, result)
		}
	}
}
