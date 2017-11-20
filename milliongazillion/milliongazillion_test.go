package milliongazillion

import "testing"

func TestStringExistsOrAdd(t *testing.T) {

	testCases := []struct {
		URL    string
		Result bool
	}{
		{"www.google.com/maps", false},
		{"www.google.com", false},
		{"www.google.com", true},
		{"www.golang.org", false},
		{"slack.padevwan.com", false}}

	trie := NewTrie()

	for idx, c := range testCases {
		if result, err := trie.StringExistsOrAdd(c.URL); err != nil {
			t.Fatalf("Error on URL %+q: %v", c.URL, err)
		} else if result != c.Result {
			t.Errorf(`Unexpected result for index %v URL %+q. Expected: %v,
				received: %v`,
				idx, c.URL, c.Result, result)
		}
	}
}
