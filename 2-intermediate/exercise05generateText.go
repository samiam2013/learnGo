package intermediate

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
	edges []*wordEdge
}

type wordEdge struct {
	to *wordNode
}

func (wN *wordNode) addEdge(next *wordNode) {
	newEdge := wordEdge{next}
	wN.edges = append(wN.edges, &newEdge)
}

func exercise05() {
	var words []*wordNode
	// you'll have to implement a sample text file if you want to run this
	f, err := os.Open("./sample.txt")
	check(err)

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
		if len(word) > 1 {
			word = strings.ToLower(word)
		}
		wordN, found = findWord(word, words)
		if found {
			lastWordNode.addEdge(wordN)
			lastWordNode = wordN
		} else {
			newWord := wordNode{word, nil}
			words = append(words, &newWord)
			lastWordNode.addEdge(&newWord)
			lastWordNode = &newWord
		}
	}
	err = scanner.Err()
	check(err)
	f.Close()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	curWord, found := findWord("the", words)
	if !found {
		curWord = words[0]
	}
	for i := 0; i < 1000; i++ {
		fmt.Print(curWord.word + " ")
		if len(curWord.edges)-1 == 0 {
			fmt.Println()
			curWord = words[r1.Intn(len(words)-1)]
		} else {
			curWord = curWord.edges[r1.Intn(len(curWord.edges)-1)].to
		}
	}
	fmt.Print("\n\n")
}

func findWord(search string, words []*wordNode) (*wordNode, bool) {
	for _, word := range words {
		if word.word == search {
			return word, true
		}
	}
	return &wordNode{"", nil}, false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
