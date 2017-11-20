package milliongazillion

//EndOfWord Key that represents the end of a word in the Children map
const EndOfWord = 0

//Node Trie node contains a map with children, where the character is the Key
type Node struct {
	Children map[byte]*Node
}

//NewNode Returns the pointer to a new instance of Node
func NewNode() *Node {
	n := new(Node)
	n.Children = make(map[byte]*Node)
	return n
}

//Trie Contains the reference to the Root node
type Trie struct {
	Root *Node
}

//NewTrie Returns a new instance of Trie
func NewTrie() *Trie {
	t := new(Trie)
	t.Root = NewNode()
	return t
}

//StringExistsOrAdd Check if a string exists in the trie, at the same time it
//adds it if it's new
func (t *Trie) StringExistsOrAdd(str string) (bool, error) {
	exists := true

	//Variable holds the current node
	node := t.Root

	//Iterate string
	for i := 0; i < len(str); i++ {

		//Byte for the current character
		char := str[i]

		//character does not exist, create a new node for that byte
		if _, nodeExists := node.Children[char]; !nodeExists {
			exists = false
			node.Children[char] = NewNode()
		}

		node = node.Children[char]
	}

	//If we are at the end of the string, but there is no EndOfWord node in there
	//children, it's a new word
	if _, nodeExists := node.Children[EndOfWord]; !nodeExists {
		node.Children[EndOfWord] = nil
		exists = false
	}

	return exists, nil
}
