package offScript

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type cell struct {
	Value    uint8
	ValueSet bool
	Possible map[int]bool
}

func (c *cell) setEmpty() {
	c.Value = 0
	c.ValueSet = false
	c.Possible = map[int]bool{
		1: true, 2: true, 3: true, 4: true, 5: true,
		6: true, 7: true, 8: true, 9: true}
}

func (c *cell) setCell(value uint8) {
	c.Value = value
	c.ValueSet = true
	c.Possible = map[int]bool{
		1: false, 2: false, 3: false, 4: false, 5: false,
		6: false, 7: false, 8: false, 9: false}
	c.Possible[int(value)] = true
}

type board struct {
	Cells [9][9]*cell
}

func (b *board) getRow(idx int) [9]*cell {
	row := [9]*cell{}
	for c := 0; c < 9; c++ {
		row[c] = b.Cells[idx][c]
	}
	return row
}

func (b *board) getCol(idx int) [9]*cell {
	col := [9]*cell{}
	for r := 0; r < 9; r++ {
		col[r] = b.Cells[r][idx]
	}
	return col
}

func (b *board) getSet(setIdx int) [9]*cell {
	set := [9]*cell{}
	colSetIdx := setIdx % 3
	rowSetIdx := (setIdx - colSetIdx) / 3
	i := 0
	for r := rowSetIdx * 3; r < (rowSetIdx*3)+3; r++ {
		for c := colSetIdx * 3; c < (colSetIdx*3)+3; c++ {
			//fmt.Printf("row %d, col %d, r %d, c %d, i %d, j %d\n",
			//	row, col, r, c, i, j)
			set[i] = b.Cells[r][c]
			i++
		}
	}
	return set
}

func (b board) rowIsValid(rowIdx int) bool {
	return cellsAreValid(b.getRow(rowIdx))
}

func (b board) colIsValid(colIdx int) bool {
	return cellsAreValid(b.getCol(colIdx))
}

func (b board) setIsValid(setIdx int) bool {
	return cellsAreValid(b.getSet(setIdx))
}

func (b board) IsValid() bool {
	for i := 0; i < 9; i++ {
		if !b.rowIsValid(i) {
			fmt.Println("row invalid, index: ", i)
			return false
		} else if !b.colIsValid(i) {
			fmt.Println("col invalid, index: ", i)
			return false
		} else if !b.setIsValid(i) {
			fmt.Println("set invalid, index: ", i)
			return false
		}
	}
	return true
}

const horizLine string = "+---------------+---------------+---------------+"
const vertLines string = "|               |               |               |"

func (b *board) Print() {
	fmt.Println(horizLine)
	fmt.Println(vertLines)
	for i := 0; i < 9; i++ {
		row := b.getRow(i)
		fmt.Print("|")
		for key, cell := range row {
			if cell != nil && cell.Value != 0 {
				fmt.Printf("  %d  ", cell.Value)
			} else {
				fmt.Print("  .  ")
			}
			if (key+1)%3 == 0 {
				fmt.Print("|")
				if key+1 == 9 {
					fmt.Println()
				}
			}
		}
		fmt.Println(vertLines)
		if (i+1)%3 == 0 && i != 8 {
			fmt.Println(horizLine)
			fmt.Println(vertLines)
		}
	}
	fmt.Println(horizLine)
}

func (b *board) rmNotPossibles() {

	// TODO move to after ingest?  // loop over each row and col and do this for set values

	for rowIdx := 0; rowIdx < 9; rowIdx++ {
		for colIdx := 0; colIdx < 9; colIdx++ {
			// only need to set not possibles for the cells with set values
			if b.Cells[rowIdx][colIdx].ValueSet {
				//fmt.Println("colIdx, rowIdx", colIdx, rowIdx)
				rowSetIdx := (rowIdx - (rowIdx % 3)) / 3
				colSetIdx := (colIdx - (colIdx % 3)) / 3
				setIdx := (rowSetIdx * 3) + (colSetIdx)
				//fmt.Println("rowset, colset, set idx:", rowSetIdx, colSetIdx, setIdx)
				valSet := b.Cells[rowIdx][colIdx].Value
				b.rmNotPossiblesCol(colIdx, valSet)
				b.rmNotPossiblesRow(rowIdx, valSet)
				b.rmNotPossiblesSet(setIdx, valSet)
			}
		}
	}
}

// tries to find hidden single, returns the number solved
func (b *board) solveHiddenSingles() int {
	// iterate over the rows and range over the columns for cells
	// to find the cells without a set value
	solvedSingles := 0
	for i := 0; i < 9; i++ {
		for _, myCell := range b.Cells[i] {
			if !myCell.ValueSet {
				// count the number of possible values
				possibleCount := 0
				lastPossible := 0
				for value, possible := range myCell.Possible {
					if possible {
						possibleCount++
						lastPossible = value
					}
				}
				if possibleCount == 1 {
					myCell.setCell(uint8(lastPossible))
					solvedSingles++
				}
			}
		}
	}
	return solvedSingles
}

func (b *board) rmNotPossiblesCol(colIdx int, value uint8) {
	col := b.getCol(colIdx)
	for i := 0; i < 9; i++ {
		if !col[i].ValueSet {
			col[i].Possible[int(value)] = false
		}
	}
	//fmt.Println("passed getCol()")
}

func (b *board) rmNotPossiblesRow(rowIdx int, value uint8) {
	row := b.getRow(rowIdx)
	for i := 0; i < 9; i++ {
		if !row[i].ValueSet {
			row[i].Possible[int(value)] = false
		}
	}
	//fmt.Println("passed getRow()")
}

func (b *board) rmNotPossiblesSet(setIdx int, value uint8) {
	set := b.getSet(setIdx)
	for i := 0; i < 9; i++ {
		if !set[i].ValueSet {
			set[i].Possible[int(value)] = false
		}
	}
}

func (b *board) Ingest(numString string) error {
	noSpaces := SpaceMap(numString)
	//log.Println(noSpaces)
	lenNoSpaces := len(noSpaces)
	if lenNoSpaces != 81 {
		log.Fatal("board.Ingest() requires 81 characters; got:", lenNoSpaces)
	}
	for i, val := range []rune(noSpaces) {
		//fmt.Printf("[%2.d]%s ", i, string(val))
		// if it's not the value of a 0 (48.(rune)) set it
		col := i % 9
		row := (i - col) / 9
		b.Cells[row][col] = &cell{}
		if val > 48 && val < 58 {
			//fmt.Printf("r%d,c%d      ", row, col)
			// subtracting the value '0' from a run gets int value from ascii
			parsedVal := uint8(val - '0')
			if !(parsedVal > 0 && parsedVal < 10) {
				return errors.New("board.Ingest() expecting value between " +
					"1 and 9; got:" + strconv.Itoa(int(parsedVal)))
			}
			b.Cells[row][col].setCell(parsedVal)
		} else if val == 48 {
			b.Cells[row][col].setEmpty()
		} else {
			return errors.New("expected rune 48-57; got: " +
				strconv.Itoa(int(val)))
		}

		// if (i+1)%9 == 0 {
		// 	fmt.Println()
		// }
	}

	b.rmNotPossibles()

	return nil
}

func cellsAreValid(cells [9]*cell) bool {
	assertOnes := map[int]int{1: 0, 2: 0, 3: 0, 4: 0,
		5: 0, 6: 0, 7: 0, 8: 0, 9: 0}
	for _, cell := range cells {
		intVal := int(cell.Value)
		assertOnes[intVal]++
		if intVal != 0 && assertOnes[intVal] > 1 {
			fmt.Println("index:", intVal, "seems to be >1 occurences")
			return false
		}
	}
	return true
}

func Run() error {
	var err error = nil
	board := board{
		Cells: [9][9]*cell{},
	}
	err = board.Ingest(EasyBoard)
	if err != nil {
		log.Fatal(err)
	}

	if false {
		zeros := map[string]interface{}{
			"row0": board.getRow(0),
			"col0": board.getCol(0),
			"set0": board.getSet(0)}
		marshalled, _ := json.MarshalIndent(zeros, "", "  ")
		fmt.Printf("%s\n", string(marshalled))
	}

	board.Print()
	fmt.Printf("is board valid?: %v\n", board.IsValid())
	solveHiddenSingles(&board)
	board.Print()
	fmt.Printf("is board valid?: %v\n", board.IsValid())

	return err

}

func solveHiddenSingles(b *board) {
	solves := b.solveHiddenSingles()
	b.rmNotPossibles()
	for solves > 0 {
		solves = b.solveHiddenSingles()
		b.rmNotPossibles()
	}
}

// SpaceMap gradually increases the amount of allocated space as more
//	non-whitespace characters are encountered
func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// credit https://stackoverflow.com/a/32081891 CC BY-SA 4.0
