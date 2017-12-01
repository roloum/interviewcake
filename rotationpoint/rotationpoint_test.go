package rotationpoint

import "testing"

//TestCases ...
var TestCases = []struct {
	List  []string
	Index int
	Err   error
}{
	{[]string{
		"ptolemaic",
		"retrograde",
		"supplant",
		"undulate",
		"xenoepist",
		"asymptote",
		"babka",
		"banoffee",
		"engender",
		"karpatka",
		"othellolagkage",
	}, 5, nil},
	{[]string{
		"flower",
		"ant",
		"banana",
		"cream",
		"donut",
		"eclipse",
	}, 1, nil},
	{[]string{
		"banana",
		"cream",
		"donut",
		"eclipse",
		"flower",
		"ant",
	}, 5, nil},
	{[]string{
		"banana",
		"cream"}, 0, nil},
	{[]string{
		"cream",
		"banana"}, 1, nil},
	{[]string{
		"cream",
		"banana",
		"mango"}, 1, nil},
	{[]string{"cream"}, 0, errMinElements},
}

func TestFindRotationPoint(t *testing.T) {

	for _, test := range TestCases {
		if index, err := FindRotationPoint(test.List); err != test.Err {
			t.Errorf("Unexpected error: %v. List: %v, expected: %v", err, test.List,
				test.Err)
		} else if index != test.Index {
			t.Errorf("Error in list %v, expected: %v, received: %v", test.List,
				test.Index, index)
		}
	}
}
