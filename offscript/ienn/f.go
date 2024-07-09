package ienn

// if err not nil (ienn)

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime/debug"
)

// F is a function that takes an error and logs it, then prints a stack trace
// and exits the program.
func F(err error) {
	if err != nil {
		log.Print(err)

		stack := debug.Stack()
		lines := bytes.Split(stack, []byte{'\n'})
		indented := make([]byte, len(stack)+len(lines))
		for _, line := range lines {
			indented = append(indented, append([]byte("\t"),
				append(line, []byte("\n")...)...)...)
		}
		fmt.Printf("Stack Trace:\n%s\n", indented)

		os.Exit(1)
	}
}
