package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testutils"
)

func TestGreetName(t *testing.T) {
	got := testutils.CaptureOutput(GreetName, "Name\n")
	if got != "What is your name? \nNice to meet you Name" {
		t.Errorf("TestGreetName() = '%s'; want 'What is your name?\nNice to meet you Name'", got)
	}
}
