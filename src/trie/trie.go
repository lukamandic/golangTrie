package trie

import (
	"fmt"
)

type Node struct {
	Value string
	Children [27]*Node
	EndOfTheWord bool
}

type Trie struct {
	Root *Node
}

func checkIfExists() {
	fmt.Println("Check if exists")
}

func (trie *Trie) SearchWord() {
	
}

func (trie *Trie) InsertNode(letter string, position int) {
	checkIfExists()
	if trie.Root == nil {
		rootNode := &Node{}
		trie.Root = rootNode
	} /*else {
		for _, element := range trie.Root.Children {
			if element != nil && if element == letter {
				
			} 
		}
	} */
}

func Print() {
	fmt.Println("From trie package")
}