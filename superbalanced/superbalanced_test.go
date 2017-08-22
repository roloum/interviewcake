package main

import "testing"

func TestBalancedTree(t *testing.T) {

	var n1, n2, n3, n4, n5, n6 = New(1), New(2), New(3), New(4), New(5), New(6)
	n3.Left = n6
	n2.Left, n2.Right = n4, n5
	n1.Left, n1.Right = n2, n3

	if IsTreeSuperBalanced(n1) == false {
		t.Error("Tree is balanced but function returned false")
	}
}

func TestUnbalancedTree(t *testing.T) {

	var n1, n2, n3, n4, n5, n6 = New(1), New(2), New(3), New(4), New(5), New(6)
	n2.Left, n4.Left, n5.Left = n4, n5, n6
	n1.Left, n1.Right = n2, n3

	if IsTreeSuperBalanced(n1) == true {
		t.Error("Tree is unbalanced but function returned true")
	}
}
