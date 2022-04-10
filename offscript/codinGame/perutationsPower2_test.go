package codingame

import (
	"reflect"
	"testing"
)

func Test_getFormulas(t *testing.T) {
	type args struct {
		varName string
		numVars int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test case",
			args: args{
				varName: "x",
				numVars: 2,
			},
			want: []string{"2**2", "2**x", "2+2", "2+x", "2-2", "2-x", "x**2", "x**x", "x+2", "x+x", "x-2", "x-x"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFormulas(tt.args.varName, tt.args.numVars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFormulas() = %v, want %v", got, tt.want)
			}
		})
	}
}
