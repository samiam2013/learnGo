package listsAndStrings

import (
	"fmt"
	"strings"
)

// Write function that translates a text to Pig Latin and back. English is
//  translated to Pig Latin by taking the first letter of every word,
//  moving it to the end of the word and adding ‘ay’.
// “The quick brown fox” becomes “Hetay uickqay rownbay oxfay”.
func exercise20() {
	sentence := "The quick brown fox jumped over the lazy dog"
	fmt.Println(convToPig(sentence))
	fmt.Println(convToEng(convToPig(sentence)))
}

func convToPig(sentence string) string {
	words := strings.Fields(sentence)
	newSentence := ""
	for i, word := range words {
		newWord := strings.ToLower(word[1:] + word[0:1] + "ay")
		if i == 0 {
			// capitalizes the first word of the translated sentence
			newWord = strings.Title(newWord)
		}
		newSentence += newWord + " "
	}
	return newSentence
}

func convToEng(sentence string) string {
	words := strings.Fields(sentence)
	newSentence := ""
	for i, word := range words {
		newWord := strings.ToLower(
			word[len(word)-3:len(word)-2] + word[:len(word)-3])
		if i == 0 {
			// capitalizes the first word of the translated sentence
			newWord = strings.Title(newWord)
		}
		newSentence += newWord + " "
	}
	return newSentence
}
