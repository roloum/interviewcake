package prodint

//GetProductsOfAllIntsExceptAtIndex receives a list of integers and returns
//the list of the products of every integer except the one at that index
func GetProductsOfAllIntsExceptAtIndex(numbers []int) []int {

	length := len(numbers)
	if length < 2 {
		panic("Array must contain at least 2 elements")
	}

	products := make([]int, length)

	//iterate the array and save in result the product of all the numbers after current index
	productSoFar := 1
	for i, val := range numbers {
		products[i] = productSoFar
		productSoFar *= val
	}

	//reverse the array multiplying the result by the product of all numbers before current index
	productSoFar = 1
	for i := len(numbers) - 1; i >= 0; i-- {
		products[i] *= productSoFar
		productSoFar *= numbers[i]
	}

	return products
}
