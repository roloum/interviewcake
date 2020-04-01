package cafeorders

//IsFirstComeFirstServed Check if we're serving orders first-come, first-served
func IsFirstComeFirstServed(takeOutOrders, dineInOrders,
	servedOrders []int) bool {

	var i, j int

	ilen, jlen, klen := len(takeOutOrders), len(dineInOrders), len(servedOrders)

	//Check that the number of orders add up
	if ilen+jlen != klen {
		return false
	}

	//Iterate the served orders as I find the order number on either the takeOut
	//or the dineIn array. Once found, increment the counter for takeOut or dineIn
	//If one server order is not found in the next position of takeOut/dineIn
	//that means the servedOrders are not first come first served
	for k := 0; k < klen; k++ {

		if i < ilen && takeOutOrders[i] == servedOrders[k] {
			i++
		} else if j < jlen && dineInOrders[j] == servedOrders[k] {
			j++
		} else {
			return false
		}

	}

	return true

}
