package intermediate

import "fmt"

/*
  We have a bad guy! He's in our system but for some reason we don't have
  hardware access ?!?! Someone told us their password is /^[a-zA-Z0-9]{8,15}$/
  and the password in alphabetical order, so now we need to guess all possible
*/

// GuessPasswords implements the interview question I couldn't answer.
func GuessPasswords() {
	fmt.Println("starting from 0 ?")
	guesses([]byte("000000000"), 0, 49, 122)
}

func guesses(guess []byte, depth int, minVal, maxVal byte) {
	for i := minVal; i <= maxVal; i++ {
		if i == 58 {
			i = 65
		} else if i == 91 {
			i = 97
		}
		guess[depth] = i
		if testPass(string(guess)) {
			fmt.Println("found it!")
			return
		} else if string(guess[3:9]) == "ABCabc" {
			fmt.Printf("wrong c: %+s\n", string(guess))
		}
		if depth < len(guess)-1 {
			guesses(guess, depth+1, i, maxVal)
		}
	}

}

func testPass(input string) bool {
	return input == "999ABCabc"
}
