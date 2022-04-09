package main

import "testing"

func Test_countLengths(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		// TODO: Add test cases.
		{
			name:   "happy path",
			input:  "#### ##     ###### ##",
			expect: "4 2 6 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if got = countLengths(tt.input); got != tt.expect {
				t.Fatalf("countLenghts(%s) = '%s'; expected '%s'", tt.input, got, tt.expect)
			}
		})
	}
}
