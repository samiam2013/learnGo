package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testUtils"
)

func TestGreetName(t *testing.T) {
	got := testUtils.CaptureOutput(GreetName, "Name\n")
	if got != "What is your name? \nNice to meet you Name" {
		t.Errorf("TestGreetName() = '%s'; want 'What is your name?\nNice to meet you Name'", got)
	}
}
