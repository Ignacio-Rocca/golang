package tries

// Node is a struct that represents a node inside the trie structure
type Node struct {
	children map[rune]*Node
	isEnd bool
}

// Trie
type Trie struct {
	root *Node

}

// InitTrie initialize an empty trie
func InitTrie() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}

// Insert
func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, r := range w {
		if currentNode.children[r] == nil {
			currentNode.children[r] = &Node{children: make(map[rune]*Node)}
		}
		currentNode = currentNode.children[r]
	}
	currentNode.isEnd = true

}

// Search
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, r := range w {
		if currentNode.children[r] == nil {
			return false
		}
		currentNode = currentNode.children[r]
	}

	return currentNode.isEnd
}