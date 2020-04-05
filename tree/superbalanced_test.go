package tree

import (
	"testing"
)

func TestFullTree(t *testing.T) {
	tree := NewNode(5)
	left := tree.InsertLeft(8)
	right := tree.InsertRight(6)
	left.InsertLeft(1)
	left.InsertRight(2)
	right.InsertLeft(3)
	right.InsertRight(4)
	result := IsBalanced(tree)
	if !result {
		t.Errorf("Received %v", result)
	}

}

func TestBothLeavesAtTheSameDepth(t *testing.T) {
	tree := NewNode(3)
	left := tree.InsertLeft(4)
	right := tree.InsertRight(2)
	left.InsertLeft(1)
	right.InsertRight(9)
	result := IsBalanced(tree)
	if !result {
		t.Errorf("Received %v", result)
	}

}

func TestLeafHeightsDifferByOne(t *testing.T) {
	tree := NewNode(6)
	tree.InsertLeft(1)
	right := tree.InsertRight(0)
	right.InsertRight(7)
	result := IsBalanced(tree)
	if !result {
		t.Errorf("Received %v", result)
	}

}

func TestLeafHeightsDifferByTwo(t *testing.T) {
	tree := NewNode(6)
	tree.InsertLeft(1)
	right := tree.InsertRight(0)
	rightRight := right.InsertRight(7)
	rightRight.InsertRight(8)
	result := IsBalanced(tree)
	if result {
		t.Errorf("Received %v", result)
	}

}

func TestThreeLeavesTotal(t *testing.T) {
	tree := NewNode(1)
	tree.InsertLeft(5)
	right := tree.InsertRight(9)
	right.InsertLeft(8)
	right.InsertRight(5)
	result := IsBalanced(tree)
	if !result {
		t.Errorf("Received %v", result)
	}

}

func TestBothSubtreesSuperbalanced(t *testing.T) {
	tree := NewNode(1)
	tree.InsertLeft(5)
	right := tree.InsertRight(9)
	rightLeft := right.InsertLeft(8)
	right.InsertRight(5)
	rightLeft.InsertLeft(7)
	result := IsBalanced(tree)
	if result {
		t.Errorf("Received %v", result)
	}

}

func TestBothSubtreesSuperbalancedTwo(t *testing.T) {
	tree := NewNode(1)
	left := tree.InsertLeft(2)
	right := tree.InsertRight(4)
	left.InsertLeft(3)
	leftRight := left.InsertRight(7)
	leftRight.InsertRight(8)
	rightRight := right.InsertRight(5)
	rightRightRight := rightRight.InsertRight(6)
	rightRightRight.InsertRight(9)
	result := IsBalanced(tree)
	if result {
		t.Errorf("Received %v", result)
	}

}

func TestOnlyOneNode(t *testing.T) {
	tree := NewNode(1)
	result := IsBalanced(tree)
	if !result {
		t.Errorf("Received %v", result)
	}

}

func TestLinkedListTree(t *testing.T) {
	tree := NewNode(1)
	right := tree.InsertRight(2)
	rightRight := right.InsertRight(3)
	rightRight.InsertRight(4)
	result := IsBalanced(tree)
	if !result {
		t.Errorf("Received %v", result)
	}
}
