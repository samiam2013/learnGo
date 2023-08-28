package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/samiam2013/learnGo/offscript/badgraph"
	"golang.org/x/term"
)

func main() {
	// check if we're running in a terminal
	if !isatty.IsTerminal(os.Stdout.Fd()) {
		log.Println("Not a terminal")
		return
	}

	// get the width of the terminal
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		log.Println(err)
		return
	}

	// print the width and height
	fmt.Printf("width: %d, height: %d\n", w, h)

	badgraph.PrintEmpty(w, h)

}
