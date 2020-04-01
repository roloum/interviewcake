package cafeorders

//IsFirstComeFirstServed Check if we're serving orders first-come, first-served
func IsFirstComeFirstServed(takeOutOrders, dineInOrders,
	servedOrders []int) bool {

	var i, j int

	ilen, jlen, klen := len(takeOutOrders), len(dineInOrders), len(servedOrders)

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

	//Did not get until the end of each array? Served orders not in order
	if i != ilen || j != jlen {
		return false
	}

	return true

}
