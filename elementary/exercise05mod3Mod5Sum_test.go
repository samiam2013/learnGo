package elementary

import (
	"testing"

	"github.com/samiam2013/learnGo/testutils"
)

// FizzBuzzSum tests Ex05 FizzBuzzSum
func TestFizzBuzzSum(t *testing.T) {
	got := testutils.CaptureOutput(FizzBuzzSum, "10\n")
	if got != "Input a nubmer: mod 3 + mod 5 sum: 33\n" {
		t.Errorf("FizzBuzzSum() = '%s'; want 'Input a nubmer: mod 3 + mod 5 sum: 33\\n", got)
	}

}
