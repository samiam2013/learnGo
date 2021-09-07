package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testutils"
)

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
