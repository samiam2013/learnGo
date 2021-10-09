package elementary

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// Write a guessing game where the user has to guess a secret number.
//  After every guess the program tells the user whether their number
//  was too large or too small. At the end the number of tries needed
//  should be printed. It counts only as one try if they input the
//  same number multiple times consecutively.

// need to test this, so I need to de-couple parts

type gameState struct {
	Cmp      int8
	Win      bool
	Response string
	Target   int64
	Guesses  map[int64]int64
}

// GuessingGame implements the exercise 9 prompt
func GuessingGame() {

	state := gameState{
		Cmp:      1,
		Win:      false,
		Response: "nothing played yet",
		Target:   -1,
		Guesses:  make(map[int64]int64)}

	guesser := state.getGuesser()

	fmt.Print("What is your guess for the random number?: ")
	var input string
	fmt.Scanln(&input)
	guess, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	guesser(guess)
	for !state.Win {
		fmt.Println(state.Response)

		fmt.Scanln(&input)
		guess, err = strconv.ParseInt(input, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		// send the new guess in and mutate the game state
		guesser(guess)
	}
	fmt.Printf("%+v\n", state)
	fmt.Println("your guess was correct! took", len(state.Guesses), "tries.")
}

func (state *gameState) getGuesser() func(int64) {
	rand.Seed(time.Now().UnixNano())
	randNum := int64(rand.Intn(1000))
	state.Target = randNum
	return func(guess int64) {
		val := state.Guesses[guess]
		if val == 0 {
			state.Guesses[guess] = 1
		} else {
			state.Guesses[guess] = val + 1
		}
		if guess < randNum {
			state.Response = "higher!: "
			state.Win = false
			state.Cmp = 1
			return
		} else if guess > randNum {
			state.Response = "lower!: "
			state.Win = false
			state.Cmp = -1
			return
		}
		state.Response = "you win!"
		state.Win = true
		state.Cmp = 0
	}
}
