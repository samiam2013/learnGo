package offScript

import "testing"

// EasyBoard _

func TestIngest(t *testing.T) {
	board := board{
		Cells: [9][9]*cell{},
	}
	err := board.Ingest(EasyBoard)
	if err != nil {
		t.Fatal(err)
	}

}
