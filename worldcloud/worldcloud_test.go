package worldcloud

import (
	"reflect"
	"testing"
)

func TestWordsToCounts(t *testing.T) {

	cases := []struct {
		s string
		m map[string]int
	}{
		//{"   a random Random! string-", map[string]int{
		{"   a Random -memory calf random! random Calf it's", map[string]int{
			"a":      1,
			"calf":   2,
			"it's":   1,
			"memory": 1,
			"random": 3}},
		{"I like cake", map[string]int{
			"I":    1,
			"like": 1,
			"cake": 1}},
		{"Chocolate cake for dinner and pound cake for dessert", map[string]int{
			"and":       1,
			"pound":     1,
			"for":       2,
			"dessert":   1,
			"Chocolate": 1,
			"dinner":    1,
			"cake":      2}},
		{"Strawberry short cake? Yum!", map[string]int{
			"cake":       1,
			"Strawberry": 1,
			"short":      1,
			"Yum":        1}},
		{"Dessert - mille-feuille cake", map[string]int{
			"cake":          1,
			"Dessert":       1,
			"mille-feuille": 1}},
		{"Mmm...mmm...decisions...decisions", map[string]int{
			"mmm":       2,
			"decisions": 2}},
		{"Allie's Bakery: Sasha's Cakes", map[string]int{
			"Bakery":  1,
			"Cakes":   1,
			"Allie's": 1,
			"Sasha's": 1}},
	}

	for row, c := range cases {
		worldCloudData := &Data{input: c.s}
		result := worldCloudData.WordsToCounts()
		if !reflect.DeepEqual(result, c.m) {
			t.Errorf("Row: %v, String: %v, Expected: %v, Received: %v", row+1, c.s,
				c.m, result)
		}
	}

}
