package main

import "testing"

func Test_mutateWithGoroutine(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "this example was put to experiment with the --race flag",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutateWithGoroutine()
		})
	}
}
