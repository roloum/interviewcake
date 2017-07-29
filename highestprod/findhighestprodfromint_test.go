package highestprod

import "testing"

func TestFindHighestProductFrom3Ints(t *testing.T) {
	cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 10, -5, 1, -100}, 5000},
		{[]int{0, 1, 1}, 0},
		{[]int{0, 2, 4, 6}, 48},
		{[]int{-3, -5, -7, -2, 0}, 0},
		{[]int{-3, -5, -7, -2, 1}, 35},
		{[]int{-3, -5, -7, -2}, -30},
	}
	for _, c := range cases {
		result := FindHighestProductFrom3Ints(c.input)
		if result != c.output {
			t.Errorf("Error on input %v, expected: %v, received: %v", c.input, c.output, result)
			break
		}
	}
}

func TestFindHighestProductFrom3IntsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("FindHighestProductFrom3Ints function did not panic")
		}
	}()
	output := FindHighestProductFrom3Ints([]int{1})
	t.Log(output)
}

func TestFindHighestProductFromKInts(t *testing.T) {
	cases := []struct {
		numbers   []int
		k, output int
	}{
		{[]int{1, 10, -5, 1, -100}, 3, 5000},
		{[]int{2, 4, -2, 7, -3, 6, 5, 9, 11}, 5, 20790},
		{[]int{2, 4, -2, 7, -3, 6, 5, 9, -21}, 5, 23814},
		{[]int{0, 1, 1}, 3, 0},
		{[]int{0, 2, 4, 6}, 3, 48},
		{[]int{-3, -5, -7, -2, 0}, 3, 0},
		{[]int{-3, -5, -7, -2, 1}, 3, 35},
		{[]int{-3, -5, -7, -2}, 3, -30},
	}
	for _, c := range cases {
		result := FindHighestProductFromKInts(c.numbers, c.k)
		if result != c.output {
			t.Errorf("Error on numbers %v, k: %v, expected: %v, received: %v", c.numbers, c.k, c.output, result)
			break
		}
	}
}

func TestFindHighestProductFromKIntsMinProdPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("FindHighestProductFromKInts function did not panic when k=2")
		}
	}()
	output := FindHighestProductFromKInts([]int{1, 1}, 2)
	t.Log(output)
}

func TestFindHighestProductFromKIntsArraySmallerThanKPanic(t *testing.T) {
	defer func() {
		if p := recover(); p == nil {
			t.Error("FindHighestProductFromKInts function did not panic when k < array elements")
		}
	}()
	output := FindHighestProductFromKInts([]int{1, 1}, 3)
	t.Log(output)
}
