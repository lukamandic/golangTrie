package trieTypes

type Node struct {
	Value string
	Children [27]*Node
	EndOfTheWord bool
}

type Trie struct {
	Root *Node
}