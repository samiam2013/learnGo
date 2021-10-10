package offscript

import (
	"encoding/json"
	"fmt"
	"log"
)

// RunSudoku is used for calling from ../main.go to test things by hand
func RunSudoku() error {
	board := board{
		Cells: [9][9]*cell{},
	}
	err := board.Ingest(hiddenSinglesBoard)
	if err != nil {
		log.Fatal(err)
	}

	if false {
		zeros := map[string]interface{}{
			"row0": getRow(&board, 0),
			"col0": getCol(&board, 0),
			"set0": getSet(&board, 0)}
		marshalled, _ := json.MarshalIndent(zeros, "", "  ")
		fmt.Printf("%s\n", string(marshalled))
	}

	board.Print()
	fmt.Printf("is board valid?: %v\n", board.IsValid())
	solvedCells := SolveNakedSingles(&board)
	board.Print()
	fmt.Println("cells solved by naked Singles:", solvedCells)
	fmt.Printf("is board valid?: %v\n", board.IsValid())

	board.Ingest(hiddenSinglesBoard)
	board.Print()
	fmt.Printf("is board valid?: %v\n", board.IsValid())
	solvedCells = SolveHiddenSingles(&board)
	board.Print()
	fmt.Println("cells solved by naked Singles:", solvedCells)
	fmt.Printf("is board valid?: %v\n", board.IsValid())

	return err

}
