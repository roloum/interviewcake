package prodint

import "testing"

func TestGetProductsOfAllIntsExceptAtIndex(t *testing.T) {
	cases := []struct {
		input, output []int
	}{
		{[]int{1, 7, 3, 4}, []int{84, 12, 28, 21}},
		{[]int{5, 3, 0}, []int{0, 0, 15}},
		{[]int{-5, 3, -1}, []int{-3, 5, -15}},
	}

	for _, tc := range cases {
		products := GetProductsOfAllIntsExceptAtIndex(tc.input)
		if len(products) == 0 {
			t.Errorf("Error on input %v, result is empty", tc.input)
		}
		for i, prod := range products {
			if prod != tc.output[i] {
				t.Errorf("Error on input %v, idx: %v. Expected: %v, received: %v", tc.input, i, tc.output[i], prod)
				break
			}
		}
	}
}

func TestGetProductsOfAllIntsExceptAtIndexPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error("GetProductsOfAllIntsExceptAtIndex did not panic")
		}
	}()

	products := GetProductsOfAllIntsExceptAtIndex([]int{1})
	t.Log(products)
}
