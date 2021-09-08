package elementary

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/samiam2013/learnGo/testutils"
)

// TestHelloWorld tests Ex01 HelloWorld (exercise01helloWorld.go)
func TestHelloWorld(t *testing.T) {
	got := testutils.CaptureOutput(HelloWorld, "")
	if got != "Hello World!\n" {
		t.Errorf("HelloWorld() = %s; want \"Hello World!\\n\"", got)
	}
}

// TestGreetName tests Ex02 GreetName (exercise02yourNameGreeting.go)
func TestGreetName(t *testing.T) {
	got := testutils.CaptureOutput(GreetName, "Name\n")
	if got != "What is your name? \nNice to meet you Name" {
		t.Errorf("TestGreetName() = '%s'; want 'What is your name?\nNice to meet you Name'", got)
	}
}

// TestGreetName tests Ex02 GreetName
func TestGreetBobOrAlice(t *testing.T) {
	got := testutils.CaptureOutput(GreetBobOrAlice, "Bob\n")
	if got != "What is your name?: Hello Bob!\n" {
		t.Errorf("TestGreetBobOrAlice() = '%s'; want 'What is your name?: Hello Bob!", got)
	}

	got = testutils.CaptureOutput(GreetBobOrAlice, "Alice\n")
	if got != "What is your name?: Hello Alice!\n" {
		t.Errorf("TestGreetBobOrAlice() = '%s'; want 'What is your name?: Hello Alice!", got)
	}

	got = testutils.CaptureOutput(GreetBobOrAlice, "Uknown\n")
	if got != "What is your name?: I don't recognize that name.\n" {
		t.Errorf("TestGreetBobOrAlice() = '%s'; want 'What is your name?: Hello Alice!", got)
	}
}

// TestTriangleSum tests Ex02 GreetName
func TestTriangleSum(t *testing.T) {
	got := testutils.CaptureOutput(TriangleSum, "10\n")
	if got != "Pick a number, an number that doesn't overflow, of course: 10?: 55\n" {
		t.Errorf("TriangleSum() = '%s'; want 'Pick a number, an number that doesn't overflow, of course: 10?: 55\\n'", got)
	}

	got = testutils.CaptureOutput(TriangleSum, "123456789\n")
	if got != "Pick a number, an number that doesn't overflow, of course: 123456789?: 7620789436823655\n" {
		t.Errorf("TriangleSum() = '%s'; want 'Pick a number, an number that doesn't overflow, of course: 10?: 55\\n'", got)
	}
}

// FizzBuzzSum tests Ex05 FizzBuzzSum
func TestFizzBuzzSum(t *testing.T) {
	got := testutils.CaptureOutput(FizzBuzzSum, "10\n")
	if got != "Input a nubmer: mod 3 + mod 5 sum: 33\n" {
		t.Errorf("FizzBuzzSum() = '%s'; want 'Input a nubmer: mod 3 + mod 5 sum: 33\\n", got)
	}

}

// CaptureOutput takes in a function to catch the output, optionally taking in lines of input for stdin
func CaptureOutput(f func(), input string) string {
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
