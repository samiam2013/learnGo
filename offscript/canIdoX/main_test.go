package main

import "testing"

func Test_mutateWithGoroutine(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "this is the barest emptiest test case ever",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mutateWithGoroutine()
		})
	}
}
