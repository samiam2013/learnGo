package offscript

import "testing"

func TestHiddenSingles(t *testing.T) {
	puzzle := board{}
	err := puzzle.Ingest(hiddenSinglesBoard)
	if err != nil {
		t.Fatal(err)
	}

	solved := board{}
	err = solved.Ingest(hiddenSinglesBoardSolved)
	if err != nil {
		t.Fatal(err)
	}

	solveHiddenSingles(&puzzle)

	if !puzzle.IsValid() {
		t.Fatal("solution invalid!")
	}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if puzzle.Cells[row][col].Value != solved.Cells[row][col].Value {
				t.Fatal("solution values didn't match!")
			}
		}
	}

}
