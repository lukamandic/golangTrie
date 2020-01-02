package main

import (
	"trie"
	"fmt"
	"strings"
	"net/http"
	//"io"
	"io/ioutil"
	//"os"
	//"bufio"
)

func processString(someWord string) {
	//fmt.Println(len(someWord))
	newTrie := &trie.Trie{Root: &trie.Node{}}
	trie.Print()
	position := 1
	for _, letter := range strings.Split(someWord, "") {
		newTrie.InsertNode(letter, position)
		position++
	}
	fmt.Println(someWord)
}

func Search(w http.ResponseWriter, r *http.Request) {
	processString(r.URL.Query()["term"][0])
}

func main() {
	dat, err := ioutil.ReadFile("./list.txt")
	if err == nil {
		countries := string(dat)
		for _, word := range countries {
			fmt.Println(string(word))
		}
	}
	http.HandleFunc("/search", Search)
	http.ListenAndServe(":8080", nil)
}