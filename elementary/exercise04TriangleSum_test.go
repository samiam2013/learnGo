package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testutils"
)

// TestTriangleSum tests Ex02 GreetName
func TestTriangleSum(t *testing.T) {
	got := testutils.CaptureOutput(TriangleSum, "10\n")
	if got != "Pick a number, an number that doesn't overflow, of course: 10?: 55\n" {
		t.Errorf("TriangleSum() = '%s'; want 'Pick a number, an number that doesn't overflow, of course: 10?: 55\\n", got)
	}

	got = testutils.CaptureOutput(TriangleSum, "123456789\n")
	if got != "Pick a number, an number that doesn't overflow, of course: 123456789?: 7620789436823655\n" {
		t.Errorf("TriangleSum() = '%s'; want 'Pick a number, an number that doesn't overflow, of course: 10?: 55\\n", got)
	}
}
