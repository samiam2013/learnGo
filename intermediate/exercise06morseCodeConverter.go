package intermediate

import (
	"fmt"
	"strings"
)

// lifted from alwindoss, thanks alwindoss.
// https://github.com/alwindoss/morse/blob/master/constants.go
var letterMorseMap = map[string]string{
	"A": ".-", "B": "-...", "C": "-.-.", "D": "-..", "E": ".", "F": "..-.",
	"G": "--.", "H": "....", "I": "..", "J": ".---", "K": "-.-", "L": ".-..",
	"M": "--", "N": "-.", "O": "---", "P": ".--.", "Q": "--.-", "R": ".-.",
	"S": "...", "T": "-", "U": "..-", "V": "...-", "W": ".--", "X": "-..-",
	"Y": "-.--", "Z": "--..", "1": ".----", "2": "..---", "3": "...--",
	"4": "....-", "5": ".....", "6": "-....", "7": "--...", "8": "---..",
	"9": "----.", "0": "-----", ".": ".-.-.-", ":": "---...", ",": "--..--",
	";": "-.-.-", "?": "..--..", "=": "-...-", "'": ".----.", "/": "-..-.",
	"!": "-.-.--", "-": "-....-", "_": "..--.-", "\"": ".-..-.", "(": "-.--.",
	")": "-.--.-", "()": "-.--.-", "$": "...-..-", "&": ".-...",
	"@": ".--.-.", "+": ".-.-."}

var morseLetterMap = map[string]string{}

func initMorseLetterMap() {
	// define the reverse map
	morseLetterMap = map[string]string{}
	for k, v := range letterMorseMap {
		morseLetterMap[v] = k
	}
}

// Write a program that automatically converts English text to Morse code
//  and vice versa.
func exercise06() {
	initMorseLetterMap()
	printMorse("Hello World!", false)
	printText(".... . .-.. .-.. ---    .-- --- .-. .-.. -.. -.-.-- ", false)
	printText(printMorse("cq cq cq w9usi cq cq cq w9usi", true), false)
}

func printMorse(text string, doReturn bool) string {
	var output string
	for i := 0; i < len(text); i++ {
		if string(text[i]) == " " {
			output += "   "
		} else {
			letter := strings.ToUpper(string(text[i]))
			output += letterMorseMap[letter] + " "
		}
	}
	if doReturn {
		return output
	}
	fmt.Println(output)
	return ""
}

func printText(morse string, doReturn bool) string {
	var morseLetter, output string
	for i := 0; i < len(morse); i++ {
		if string(morse[i]) != " " {
			morseLetter += string(morse[i])
		} else {
			output += morseLetterMap[morseLetter]
			morseLetter = ""
			if len(morse)-2 > i {
				if string(morse[i+1])+string(morse[i+2]) == "  " {
					output += " "
					i += 2
				}
			}
		}
	}
	if doReturn {
		return output
	}
	fmt.Println(output)
	return ""
}
