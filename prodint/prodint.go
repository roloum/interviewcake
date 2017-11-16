package prodint

//GetProductsOfAllIntsExceptAtIndex receives a list of integers and returns
//the list of the products of every integer except the one at that index
func GetProductsOfAllIntsExceptAtIndex(numbers []int) []int {

	numbersLen := len(numbers)

	if numbersLen < 2 {
		panic("Not enough integers")
	}

	//Iterate the array twice to compute the products before and after the index
	//Compute the productSoFar as we iterate the array before index
	//First iteration initializes the result
	productSoFar := 1
	products := make([]int, numbersLen)
	for i := 0; i < numbersLen; i++ {
		products[i] = productSoFar
		productSoFar *= numbers[i]
	}

	//Second iteration reverses the array and computes the productSoFar after index
	//and multiply the current result in product by productSoFar
	productSoFar = 1
	for i := numbersLen - 1; i >= 0; i-- {
		products[i] *= productSoFar
		productSoFar *= numbers[i]
	}

	return products
}
