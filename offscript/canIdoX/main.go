package main

import (
	"log"
	"time"
)

// can I log from a goroutine?
func logFromGoroutineWithCallback() {
	c := make(chan func())
	go func(callback chan func()) {
		callback <- func() {
			log.Println("log called from func() returned over a channel by a go routine")
		}
	}(c)
	(<-c)() // this doesn't look like anything but it calls the returned func()
}

func logFromGoroutine() {
	go func() {
		log.Println("this is a log print called from a go routine")
	}()
}

func panicFromGoroutine() {
	go func() {
		panic("this is a panic from a goroutine.")
	}()
}

var stringToMutate = "Hello World"

func mutateWithGoroutine() {
	log.Println("before goroutine mutation:", stringToMutate)
	// open a bool channel to block for goroutine to finish
	//	when waiting to check value again
	c := make(chan bool)
	go func(str *string, c chan bool) {
		*str = "changed value"
		c <- true
	}(&stringToMutate, c)
	_ = <-c // blocking here

	log.Println("after goroutine mutation:", stringToMutate)

}

func main() {
	// can I ...
	logFromGoroutine()
	// no you can't. not unless you twiddle your thumbs and wait for it
	// or create a channel to return it

	// can I ...
	logFromGoroutineWithCallback()
	// yes you can.

	// can I ...
	if false {
		panicFromGoroutine()
		// not necessarily.
		// rather you can but you have to wait for it
		// this is a stupid hack, you should waid for a signal on a channel instead
		// but that sort of defeats the goal of simplest way to propogate failure
		time.Sleep(time.Second * 1)
	}

	// can I mutate with goroutine?
	mutateWithGoroutine()
	// yes you can

	// does a test for that notice a data race possibility?
	// it doesn't appear to

	// can I ...
	// ?
}
