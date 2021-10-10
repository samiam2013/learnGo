package offscript

import (
	"encoding/json"
	"fmt"
	"log"
)

// RunSudoku is used for calling from ../main.go to test things by hand
func RunSudoku() error {
	b := board{
		Cells: [9][9]*cell{},
	}
	err := b.Ingest(hiddenSinglesBoard)
	if err != nil {
		log.Fatal(err)
	}

	if false {
		zeros := map[string]interface{}{
			"row0": getRow(&b, 0),
			"col0": getCol(&b, 0),
			"set0": getSet(&b, 0)}
		marshalled, _ := json.MarshalIndent(zeros, "", "  ")
		fmt.Printf("%s\n", string(marshalled))
	}

	fns := []func(*board) int{SolveNakedSingles, SolveHiddenSingles}
	for _, solver := range fns {
		b.Ingest(hiddenSinglesBoard)
		solvedCells := solver(&b)
		b.Print()
		fmt.Printf("cells solved: %d - valid: %v\n", solvedCells, b.IsValid())
	}
	return err

}
