package elementary

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"
	"testing"
)

// TestHelloWorld tests Ex01 HelloWorld (exercise01helloWorld.go)
func TestHelloWorld(t *testing.T) {
	got := captureOutput(HelloWorld, "")
	if got != "Hello World!\n" {
		t.Errorf("HelloWorld() = %s; want \"Hello World!\\n\"", got)
	}
}

// TestGreetName tests Ex02 GreetName (exercise02yourNameGreeting.go)
func TestGreetName(t *testing.T) {
	got := captureOutput(GreetName, "Name\n")
	if got != "What is your name? \nNice to meet you Name" {
		t.Errorf("TestGreetName() = '%s'; want 'What is your name?\nNice to meet you Name'", got)
	}
}

// TestGreetName tests Ex02 GreetName
func TestGreetBobOrAlice(t *testing.T) {
	got := captureOutput(GreetBobOrAlice, "Bob\n")
	if got != "What is your name?: Hello Bob!\n" {
		t.Errorf("TestGreetBobOrAlice() = '%s'; want 'What is your name?: Hello Bob!", got)
	}

	got = captureOutput(GreetBobOrAlice, "Alice\n")
	if got != "What is your name?: Hello Alice!\n" {
		t.Errorf("TestGreetBobOrAlice() = '%s'; want 'What is your name?: Hello Alice!", got)
	}

	got = captureOutput(GreetBobOrAlice, "Uknown\n")
	if got != "What is your name?: I don't recognize that name.\n" {
		t.Errorf("TestGreetBobOrAlice() = '%s'; want 'What is your name?: Hello Alice!", got)
	}
}

// TestTriangleSum tests Ex02 GreetName
func TestTriangleSum(t *testing.T) {
	got := captureOutput(TriangleSum, "10\n")
	if got != "Pick a number, an number that doesn't overflow, of course: 10?: 55\n" {
		t.Errorf("TriangleSum() = '%s'; want 'Pick a number, an number that doesn't overflow, of course: 10?: 55\\n'", got)
	}

	got = captureOutput(TriangleSum, "123456789\n")
	if got != "Pick a number, an number that doesn't overflow, of course: 123456789?: 7620789436823655\n" {
		t.Errorf("TriangleSum() = '%s'; want 'Pick a number, an number that doesn't overflow, of course: 10?: 55\\n'", got)
	}
}

// FizzBuzzSum tests Ex05 FizzBuzzSum
func TestFizzBuzzSum(t *testing.T) {
	got := captureOutput(FizzBuzzSum, "10\n")
	if got != "Input a nubmer: mod 3 + mod 5 sum: 33\n" {
		t.Errorf("FizzBuzzSum() = '%s'; want 'Input a nubmer: mod 3 + mod 5 sum: 33\\n", got)
	}

}

// TestTriangleOrFac tests Ex 05
func TestTriangleOrFac(t *testing.T) {
	got := captureOutput(TriangleOrFac, "10\ns\n")
	if got != "enter a number: would you like to compute the (s)um or (p)roduct?: Your result: 55\n" {
		t.Errorf("TriangleOrFac() = '%s'; want 'enter a number: would you like to compute the "+
			"(s)um or (p)roduct?: Your result: 55\\n'", got)
	}

	got = captureOutput(TriangleOrFac, "6\np\n")
	if got != "enter a number: would you like to compute the (s)um or (p)roduct?: Your result: 720\n" {
		t.Errorf("TriangleOrFac() = '%s'; want 'enter a number: would you like to compute the "+
			"(s)um or (p)roduct?: Your result: 720\\n'", got)
	}
}

// TestMultiplicationTable tests Ex 07
func TestMultiplicationTable(t *testing.T) {
	got := captureOutput(MultiplicationTable, "")
	lastLine := strings.Split(got, "\n")[11]
	if lastLine != "  12  24  36  48  60  72  84  96 108 120 132 144" {
		t.Errorf("TriangleOrFac() = '%s'; want '  12  24  36  48  60  72  84  96 108 120 132 144'", lastLine)
	}
}

func TestCheckBigPrime(t *testing.T) {
	primesMap := map[int64]bool{
		3: true, 5: true, 7: true, 11: true, 13: true,
		17: true, 19: true, 23: true, 29: true, 31: true,
		2: false, 6: false, 8: false, 12: false, 15: false,
		18: false, 20: false, 25: false, 30: false, 32: false}
	for k, v := range primesMap {
		checkReturn := CheckBigPrime(new(big.Int).SetInt64(k))
		if checkReturn != v {
			t.Errorf("CheckBigPrime(%d) = '%v'; want %v", k, checkReturn, v)
		}
	}
}

// captureOutput takes in a function to catch the output, optionally taking in lines of input for stdin
func captureOutput(f func(), input string) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	stdin := os.Stdin
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		os.Stdin = stdin
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	if input != "" {
		content := []byte(input)
		tmpfile, err := ioutil.TempFile("", "example")
		defer os.Remove(tmpfile.Name())
		if err != nil {
			log.Fatal(err)
		}
		if _, err := tmpfile.Write(content); err != nil {
			log.Fatal(err)
		}
		if _, err := tmpfile.Seek(0, 0); err != nil {
			log.Fatal(err)
		}
		os.Stdin = tmpfile
	}

	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
