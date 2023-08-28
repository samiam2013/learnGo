package badgraph

import "strings"

type Matrix [][]bool

func (m Matrix) String() string {
	var s string
	for _, row := range m {
		for _, cell := range row {
			if cell {
				s += "█"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	return strings.TrimRight(s, "\n")
}

func PrintEmpty(w, h int) {
	m := make(Matrix, h)
	for i := range m {
		m[i] = make([]bool, w)
	}
	println(m.String())
}
