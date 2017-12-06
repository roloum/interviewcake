package cakethief

import "errors"

var errEmptyList = errors.New("List is empty")
var errNegativeValue = errors.New("Value can't be negative")
var errCeroValue = errors.New("Zero value")

//MaxDuffelBagValue Based on a list of tuples (weight, value), returns the
//maximum value for a given capacity.
//We break down the problem by figuring out the max value for each capacity
//from 0 to maxCapacity.
func MaxDuffelBagValue(list [][]int, capacity int) (int, error) {

	count := len(list)
	if count == 0 {
		return 0, errEmptyList
	}

	//Lets create an array where we store the maxValue per capacity, starting at
	//index 0 until the capacity
	values := make([]int, capacity+1)

	//Process the maxValue for each capacity from 0 to max capacity
	for i := 1; i <= capacity; i++ {

		maxValue := 0

		for _, tuple := range list {
			if tuple[0] < 0 || tuple[1] < 0 {
				return 0, errNegativeValue
			} else if tuple[0] == 0 {
				return 0, errCeroValue
			} else if tuple[0] <= i {
				//Only process tuples with weigth <= current capacity

				//Multiply quotient by value and Add the value for the remainder
				//values array should have maxValue for remainder already
				currentValue := (int(i/tuple[0]) * tuple[1]) + values[i%tuple[0]]

				//Check maxValue
				if currentValue > maxValue {
					maxValue = currentValue
				}

			}
		}

		//Assign maxValue to current capacity
		values[i] = maxValue

	}

	//Return value for maxCapacity
	return values[capacity], nil
}
