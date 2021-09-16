package elementary

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
)

// TestHelloWorld tests Ex01 HelloWorld (exercise01helloWorld.go)
func TestHelloWorld(t *testing.T) {
	got := captureOutput(HelloWorld, "")
	correct := "Hello World!\n"
	if got != correct {
		t.Errorf("HelloWorld() = '%s'; want '%s'", got, correct)
	}
}

// TestGreetName tests Ex02 GreetName (exercise02yourNameGreeting.go)
func TestGreetName(t *testing.T) {
	got := captureOutput(GreetName, "Name\n")
	correct := "What is your name? \nNice to meet you Name"
	if got != correct {
		t.Errorf("TestGreetName() = '%s'; want '%s'", got, correct)
	}
}

// TestGreetName tests Ex02 GreetName
func TestGreetBobOrAlice(t *testing.T) {
	testCases := map[string]string{
		"Bob\n":    "Hello Bob!\n",
		"Alice\n":  "Hello Alice!\n",
		"Uknown\n": "I don't recognize that name.\n",
	}
	for input, correct := range testCases {
		got := captureOutput(GreetBobOrAlice, input)
		got = got[strings.LastIndex(got, ":")+2:]
		if got != correct {
			t.Errorf("TestGreetBobOrAlice() = '%s'; want '%s'", got, correct)
		}
	}
}

// TestTriangleSum tests Ex02 GreetName
func TestTriangleSum(t *testing.T) {
	testCases := map[int64]int64{
		10:        55,
		123456789: 7620789436823655,
	}
	for input, correct := range testCases {
		got := captureOutput(TriangleSum, strconv.FormatInt(input, 10)+"\n")
		got = got[strings.LastIndex(got, ":")+2:]
		got = strings.TrimSuffix(got, "\n")
		gotNum, err := strconv.ParseInt(got, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		if gotNum != correct {
			t.Errorf("TriangleSum() = '%d'; want '%d'", gotNum, correct)
		}
	}
}

// FizzBuzzSum tests Ex05 FizzBuzzSum
func TestFizzBuzzSum(t *testing.T) {
	testCases := map[int]int64{
		0: 0, 1: 0, 2: 0, 3: 3, 4: 3,
		5: 8, 6: 14, 10: 33, 165: 6435,
		6435:    9665370,
		9665370: 21797859521295,
	}
	for input, correct := range testCases {
		got := captureOutput(FizzBuzzSum, strconv.Itoa(input)+"\n")
		got = got[strings.LastIndex(got, ":")+2:]
		got = strings.TrimSuffix(got, "\n")
		gotNum, err := strconv.ParseInt(got, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		if gotNum != correct {
			t.Errorf("FizzBuzzSum() = '%s'; want '%v'", got, correct)
		}
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
		t.Errorf("TriangleOrFac() = '%s'; want '  12  24  36  48  60  72  84"+
			" 96 108 120 132 144'", lastLine)
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

func TestGuessingGame(t *testing.T) {
	var buf bytes.Buffer
	for i := 0; i <= 1000; i++ {
		buf.WriteString(strconv.Itoa(i) + "\n")
	}
	response := captureOutput(GuessingGame, buf.String())
	last := response[strings.LastIndex(response, ":")+2:]
	congrats := last[:strings.LastIndex(last, "!")+1]
	correct := "your guess was correct!"
	if congrats != correct {
		t.Errorf("GuessingGame() = '%s'; want '%s'", congrats, correct)
	}
}

func TestFizzBuzz(t *testing.T) {
	response := FizzBuzz(15, true)
	noLastNewLine := response[:strings.LastIndex(response, "\n")]
	t.Log(noLastNewLine)
	lastLine := noLastNewLine[strings.LastIndex(noLastNewLine, "\n")+1:]
	correct := "fizzbuzz"
	if lastLine != correct {
		t.Errorf("FizzBuzz() = '%s'; want '%s'", lastLine, correct)
	}

}

func TestAnagrams(t *testing.T) {
	testCases := map[[2]string]bool{
		{"dormitory", "dirtyroom"}:         true,
		{"school master", "the classroom"}: true,
		{"conversation", "voicesranton"}:   true,
		{"listen", "silent"}:               true,
		{"astronomer", "moonstarer"}:       true,
		{"the eyes", "they see"}:           true,
		{"funeral", "realfun"}:             true,
		{"yes", "no"}:                      false,
		{"left", "right"}:                  false,
		{"hello", "goodbye"}:               false,
		{"test", "ignore"}:                 false,
	}
	for k, v := range testCases {
		response := IsAnagram(k[0], k[1])
		if response != v {
			t.Errorf("isAnagram('%s','%s') = '%v'; wanted '%v'",
				k[0], k[1], response, v)
		}
		response = IsAnagramFast(k[0], k[1])
		if response != v {
			t.Errorf("isAnagramFast('%s','%s') = '%v'; wanted '%v'",
				k[0], k[1], response, v)
		}
	}
}

func TestAlternatingSeries(t *testing.T) {
	response := AlternatingSeries()
	correct := float64(-1.8274410005639339)
	if response != correct {
		t.Errorf("AlternatingSeries() = '%v'; want '%v'", response, correct)
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
