package secondlargest

import "testing"

type TestCase struct {
	Root   *Node
	Result int
}

func TestFindSecondLargest(t *testing.T) {
	cases := []TestCase{
		{&Node{25,
			&Node{11,
				&Node{9, nil, nil},
				&Node{17, nil, nil}},
			&Node{40,
				&Node{29, nil, nil},
				nil}},
			29},
		{&Node{5,
			&Node{3,
				&Node{1, nil, nil},
				nil},
			nil},
			3},
		{&Node{30,
			&Node{15, nil, nil},
			&Node{40,
				&Node{35,
					&Node{33, nil, nil},
					nil},
				nil}},
			35}}

	for idx, c := range cases {
		if result := FindSecondLargest(c.Root); result != c.Result {
			t.Errorf("Wrong result for test case with index %d, expected %d, received %d",
				idx, c.Result, result)
		}
	}
}

func TestFindSecondLargestNotEnoughNodes(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Function did not panic when called with not not enough nodes")
		}
	}()

	cases := []TestCase{{nil, 0}, {&Node{1, nil, nil}, 0}}
	for _, c := range cases {
		_ = FindSecondLargest(c.Root)
	}

}
