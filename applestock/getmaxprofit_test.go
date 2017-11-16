package applestock

import "testing"

//Test case for GetMaxProfit
func TestGetMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		profit int
	}{
		{[]int{10, 7, 11, 5, 8, 11, 9}, 6},
		{[]int{10, 7, 11, 5, 8, 11, 17, 3, 9, 16}, 13},
		{[]int{10, 8, 7, 5, 3, 1}, -1},
		{[]int{10, 10}, 0},
	}

	for _, c := range cases {
		if maxProfit := GetMaxProfit(c.prices); maxProfit != c.profit {
			t.Errorf("Bad profit! Expected: %v, returned: %v", c.profit, maxProfit)
		}
	}
}

//Tests panic when list does not have enough prices
func TestGetMaxProfitPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error("GetMaxProfit call did not panic")
		}
	}()

	maxProfit := GetMaxProfit([]int{10})
	t.Log(maxProfit)
}

//Tests panic for negative price
func TestGetMaxProfitNegativePricePanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error("GetMaxProfit call did not panic with negative input")
		}
	}()

	maxProfit := GetMaxProfit([]int{-1, 0})
	t.Log(maxProfit)
}
