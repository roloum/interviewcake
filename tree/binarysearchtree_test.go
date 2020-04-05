package tree

import (
	"testing"
)

func TestValidFullTree(t *testing.T) {
	tree := NewNode(50)
	left := tree.InsertLeft(30)
	right := tree.InsertRight(70)
	left.InsertLeft(10)
	left.InsertRight(40)
	right.InsertLeft(60)
	right.InsertRight(80)
	result := IsBinarySearch(tree)
	if !result {
		t.Errorf("Received %v", result)
	}
}

func TestBothSubtreesValid(t *testing.T) {
	tree := NewNode(50)
	left := tree.InsertLeft(30)
	right := tree.InsertRight(80)
	left.InsertLeft(20)
	left.InsertRight(60)
	right.InsertLeft(70)
	right.InsertRight(90)
	result := IsBinarySearch(tree)
	if result {
		t.Errorf("Received %v", result)
	}
}

func TestBothSubtreesValidRight(t *testing.T) {
	tree := NewNode(50)
	left := tree.InsertLeft(30)
	right := tree.InsertRight(80)
	left.InsertLeft(20)
	left.InsertRight(40)
	right.InsertLeft(45)
	right.InsertRight(90)
	result := IsBinarySearch(tree)
	if result {
		t.Errorf("Received %v", result)
	}
}

func TestDescendingLinkedList(t *testing.T) {
	tree := NewNode(50)
	left := tree.InsertLeft(40)
	leftLeft := left.InsertLeft(30)
	leftLeftLeft := leftLeft.InsertLeft(20)
	leftLeftLeft.InsertLeft(10)
	result := IsBinarySearch(tree)
	if !result {
		t.Errorf("Received %v", result)
	}
}

func TestOutOfOrderLinkedList(t *testing.T) {
	tree := NewNode(50)
	right := tree.InsertRight(70)
	rightRight := right.InsertRight(60)
	rightRight.InsertRight(80)
	result := IsBinarySearch(tree)
	if result {
		t.Errorf("Received %v", result)
	}
}

func TestOneNodeTree(t *testing.T) {
	tree := NewNode(50)
	result := IsBinarySearch(tree)
	if !result {
		t.Errorf("Received %v", result)
	}
}
