package balancedbtree

import (
	"testing"
	"time"
)

var n9 = Node{9, nil, nil}
var n8 = Node{8, &n9, nil}
var n4 = Node{4, &n8, nil}
var n5 = Node{5, nil, nil}
var n2 = Node{2, &n4, &n5}
var n6 = Node{6, nil, nil}
var n7 = Node{7, nil, nil}
var n3 = Node{3, &n6, &n7}
var n1 = Node{1, &n2, &n3}

//TestIsTreeBalanced ...
func TestIsTreeBalanced(t *testing.T) {

	start := time.Now()
	if result := IsTreeBalanced(&n1); result != false {
		t.Fatalf("Tree w/ root n1 is unbalanced. Received %v", result)
	}
	t.Logf("Checking unbalanced tree using stack: %v", time.Since(start).Nanoseconds())

	start = time.Now()
	if result := IsTreeBalanced(&n3); result != true {
		t.Fatalf("Tree w/ root n3 is balanced. Received %v", result)
	}
	t.Logf("Checking balanced tree using stack: %v", time.Since(start).Nanoseconds())
}

//TestIsTreeBalanced ...
func TestIsTreeSuperBalancedBF(t *testing.T) {

	start := time.Now()
	if result := IsTreeSuperBalancedBF(&n1); result != false {
		t.Fatalf("Tree w/ root n1 is unbalanced. Received %v", result)
	}
	t.Logf("Checking unbalanced tree using stack: %v", time.Since(start).Nanoseconds())

	start = time.Now()
	if result := IsTreeSuperBalancedBF(&n3); result != true {
		t.Fatalf("Tree w/ root n3 is balanced. Received %v", result)
	}
	t.Logf("Checking balanced tree using stack: %v", time.Since(start).Nanoseconds())
}
