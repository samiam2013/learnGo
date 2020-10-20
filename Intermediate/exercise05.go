package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// Write a program that automatically generates essays for you.
//  (1) Using a sample text, create a directed (multi-)graph where the words of
//   a text are nodes and there is a directed edge between u and v if u is
//   followed by v in your sample text.
//   Multiple occurrences lead to multiple edges.
//  (2) Do a random walk on this graph: Starting from an arbitrary node choose
//   a random successor. If no successor exists, choose another random node.

type wordNode struct {
	word  string
	edges []wordEdge
}

func (wN *wordNode) addEdge(next *wordNode) {
	newEdge := wordEdge{next}
	wN.edges = append(wN.edges, newEdge)
}

type wordEdge struct {
	to *wordNode
}

func exercise05() {
	/*
		word1 := wordNode{"hello", nil}
		word2 := wordNode{"foo", nil}
		word3 := wordNode{"world", nil}
		word1.addEdge(&word2)
		word1.addEdge(&word3)
		fmt.Println("len word1 edges: ", len(word1.edges))
		if true {
			panic(nil)
		}
	*/

	words := *loadWords()
	wordsLen := len(words)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	word := words[0] // _ := findWord("the", words)
	for i := 0; i < 10; i++ {
		fmt.Print(word.word + " ")
		fmt.Print("len edges ", len(word.edges))
		if len(word.edges) == 0 {
			fmt.Println()
			word = words[r1.Intn(wordsLen-1)]
		} else {
			word = *word.edges[r1.Intn(len(word.edges)-1)].to
		}
	}
}

func loadWords() *[]wordNode {
	var words []wordNode
	// you'll have to implement a sample text file if you want to run this
	f, err := os.Open("./sample1.txt")
	check(err)

	// read in the file
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	reg, err := regexp.Compile("[^a-zA-Z]+")
	check(err)

	lastWordNode := &wordNode{scanner.Text(), nil}
	var found bool
	var word string
	var wordN *wordNode
	for scanner.Scan() {
		word = scanner.Text()
		word = reg.ReplaceAllString(word, "")
		word = strings.ToLower(word)
		//fmt.Println(word)
		wordN, found = findWord(word, words)
		if found {
			lastWordNode.addEdge(wordN)
			lastWordNode = wordN
			//if len(lastWordNode.edges) > 1 {
			//	fmt.Println("edges len: ", len(lastWordNode.edges))
			//}
		} else {
			newWord := wordNode{word, nil}
			words = append(words, newWord)
			lastWordNode.addEdge(&newWord)
			lastWordNode = &newWord
		}
	}
	err = scanner.Err()
	check(err)
	f.Close()
	fmt.Println("words len", len(words))
	return &words
}

func findWord(search string, words []wordNode) (*wordNode, bool) {
	for _, word := range words {
		if word.word == search {
			return &word, true
		}
	}
	return &wordNode{"", nil}, false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
