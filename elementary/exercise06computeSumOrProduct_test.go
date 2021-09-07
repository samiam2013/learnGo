package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testutils"
)

// TestTriangleOrFac tests Ex 05
func TestTriangleOrFac(t *testing.T) {
	got := testutils.CaptureOutput(TriangleOrFac, "10\ns\n")
	if got != "enter a number: would you like to compute the (s)um or (p)roduct?: Your result: 55\n" {
		t.Errorf("TriangleOrFac() = '%s'; want 'enter a number: would you like to compute the "+
			"(s)um or (p)roduct?: Your result: 55\\n'", got)
	}

	got = testutils.CaptureOutput(TriangleOrFac, "6\np\n")
	if got != "enter a number: would you like to compute the (s)um or (p)roduct?: Your result: 720\n" {
		t.Errorf("TriangleOrFac() = '%s'; want 'enter a number: would you like to compute the "+
			"(s)um or (p)roduct?: Your result: 720\\n'", got)
	}
}
