package codingame

import (
	"fmt"
	"strconv"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func countLengths(input string) (out string) {
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	var last rune = rune(input[0])
	sum := 0
	for _, r := range input {
		if r == last {
			sum++
			continue
		}
		if sum > 1 {
			out += strconv.Itoa(sum) + " "
			sum = 0
		}

	}
	out += fmt.Sprintf("%d", sum)
	return
	//fmt.Println("1 2 3")// Write answer to stdout
}
