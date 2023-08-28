package badgraph

import "testing"

func TestMatrix_String(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want string
	}{
		{"empty", [][]bool{{false, true}, {false, false}}, " █\n  "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Matrix.String() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
