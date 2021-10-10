package offscript

import "testing"

// TestNakedSingles runs a test case for a hidden singles puzzle solving
func TestNakedSingles(t *testing.T) {
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

	solvedCells := SolveNakedSingles(&puzzle)
	if solvedCells != 51 {
		t.Fatal("expected 51 cells solved, got:", solvedCells)
	}

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

// TestNakedSingles runs a test case for a hidden singles puzzle solving
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

	solvedCells := SolveHiddenSingles(&puzzle)
	if solvedCells != 51 {
		t.Fatal("expected 51 cells solved, got:", solvedCells)
	}

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
