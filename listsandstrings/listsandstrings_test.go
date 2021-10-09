package listsandstrings

import (
	"testing"
)

func TestLargestInt(t *testing.T) {
	output := LargestInt([]int32{830, 1, 58, 282, 500, 32, 84})
	if output != 830 {
		t.Fatal("largest number result wrong:", output)
	}
}
