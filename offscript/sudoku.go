package offscript

import (
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
	Possible map[uint8]bool
}

func (c *cell) setEmpty() {
	c.Value = 0
	c.ValueSet = false
	c.Possible = map[uint8]bool{
		1: true, 2: true, 3: true, 4: true, 5: true,
		6: true, 7: true, 8: true, 9: true}
}

func (c *cell) setCell(value uint8) {
	c.Value = value
	c.ValueSet = true
	c.Possible = map[uint8]bool{
		1: false, 2: false, 3: false, 4: false, 5: false,
		6: false, 7: false, 8: false, 9: false}
	c.Possible[value] = true
}

type board struct {
	Cells [9][9]*cell
}

func getRow(b *board, idx int) [9]*cell {
	row := [9]*cell{}
	for c := 0; c < 9; c++ {
		row[c] = b.Cells[idx][c]
	}
	return row
}

func getCol(b *board, idx int) [9]*cell {
	col := [9]*cell{}
	for r := 0; r < 9; r++ {
		col[r] = b.Cells[r][idx]
	}
	return col
}

func getSet(b *board, setIdx int) [9]*cell {
	set := [9]*cell{}
	colSetIdx := setIdx % 3
	rowSetIdx := (setIdx - colSetIdx) / 3
	i := 0
	for r := rowSetIdx * 3; r < (rowSetIdx*3)+3; r++ {
		for c := colSetIdx * 3; c < (colSetIdx*3)+3; c++ {
			set[i] = b.Cells[r][c]
			i++
		}
	}
	return set
}

func (b board) subSetIsValid(setIdx int,
	getter func(*board, int) [9]*cell) bool {
	return cellsAreValid(getter(&b, setIdx))
}

func (b board) IsValid() bool {
	for i := 0; i < 9; i++ {
		if !b.subSetIsValid(i, getRow) {
			fmt.Println("row invalid, index: ", i)
			return false
		} else if !b.subSetIsValid(i, getCol) {
			fmt.Println("col invalid, index: ", i)
			return false
		} else if !b.subSetIsValid(i, getSet) {
			fmt.Println("set invalid, index: ", i)
			return false
		}
	}
	return true
}

// cell horizontal
const cH string = "---------------"

// horizontal union
const hUn string = "+"

// vertical union
const vUn string = "|"

// the padding on either side of a value
const pad string = "  "

// cell vertical - described as an empty row with spaces instead of noVals
const dblPad string = pad + pad
const cV string = pad + " " + dblPad + " " + dblPad + " " + pad

const noVal string = "."

const horizLine string = hUn + cH + hUn + cH + hUn + cH + hUn
const vertLines string = vUn + cV + vUn + cV + vUn + cV + vUn

func (b *board) Print() {
	fmt.Println(horizLine)
	fmt.Println(vertLines)
	for i := 0; i < 9; i++ {
		row := getRow(b, i)
		fmt.Print(vUn)
		for key, cell := range row {
			if cell != nil && cell.Value != 0 {
				fmt.Printf(pad+"%d"+pad, cell.Value)
			} else {
				fmt.Print(pad + noVal + pad)
			}
			if (key+1)%3 == 0 {
				fmt.Print(vUn)
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

func computeSetIdx(rowIdx, colIdx int) int {
	// computer set indexes 1-9 -> 1-3 for row, col to get setIdx
	rowSetIdx := (rowIdx - (rowIdx % 3)) / 3
	colSetIdx := (colIdx - (colIdx % 3)) / 3
	return (rowSetIdx * 3) + (colSetIdx)
}

func (b *board) rmNotPossibles() {
	// loop over each row and col
	//	and remove possibles in row, col, set for set values
	for rowIdx := 0; rowIdx < 9; rowIdx++ {
		for colIdx := 0; colIdx < 9; colIdx++ {
			// only need to set not possibles for the cells with set values
			if b.Cells[rowIdx][colIdx].ValueSet {
				setIdx := computeSetIdx(rowIdx, colIdx)
				val := b.Cells[rowIdx][colIdx].Value
				b.rmNotPossiblesSet(colIdx, val, getCol)
				b.rmNotPossiblesSet(rowIdx, val, getRow)
				b.rmNotPossiblesSet(setIdx, val, getSet)
			}
		}
	}
}

// tries to find hidden single, returns the number solved
func (b *board) solveNakedSingles() int {
	// iterate over the rows and range over the columns for cells
	// to find the cells without a set value
	solvedCells := 0
	for i := 0; i < 9; i++ {
		for _, myCell := range b.Cells[i] {
			if !myCell.ValueSet {
				// count the number of possible values, keeping the last one
				possibleCount := 0
				lastPossible := uint8(0)
				for value, possible := range myCell.Possible {
					if possible {
						possibleCount++
						lastPossible = value
					}
				}
				// if there's only one set the last value for the cell
				if possibleCount == 1 {
					myCell.setCell(uint8(lastPossible))
					solvedCells++
				}
			}
		}
	}
	return solvedCells
}

// tries to find hidden single, returns the number solved
func (b *board) solveHiddenSingles(getter func(*board, int) [9]*cell) int {
	solvedCells := 0
	// for each (row|col|set)
	for idx := 0; idx < 9; idx++ {
		//fmt.Println("idx:", idx)
		// get all the cells at this index
		cells := getter(b, idx)
		var d uint8
		// for each possible value check all the cells
		for d = 1; d < 10; d++ {
			//fmt.Println("d:", d)
			// for each cell check all the other cells
			for cIdx := 0; cIdx < 9; cIdx++ {
				//fmt.Println("cIdx:", cIdx)
				myCell := cells[cIdx]
				if myCell.ValueSet {
					continue
				}
				numCands := 0
				// for each other cell
				for other := 0; other < 9; other++ {
					// check our cesll is not the other cell
					otherCell := cells[other]
					if myCell == otherCell {
						continue
					}
					//fmt.Println("other:", other)
					if otherCell.Possible[d] {
						numCands++
						//fmt.Printf("count candidates:",numCands)
					}
				}
				if numCands == 0 {
					myCell.setCell(d)
					solvedCells++
				}
			}
		}
	}
	return solvedCells
}

func (b *board) rmNotPossiblesSet(idx int, value uint8,
	getter func(*board, int) [9]*cell) {
	set := getter(b, idx)
	for i := 0; i < 9; i++ {
		if !set[i].ValueSet {
			set[i].Possible[value] = false
		}
	}
}

func (b *board) Ingest(numString string) error {
	noSpaces := spaceMap(numString)
	lenNoSpaces := len(noSpaces)
	if lenNoSpaces != 81 {
		log.Fatal("board.Ingest() requires 81 characters; got:", lenNoSpaces)
	}
	for i, val := range []rune(noSpaces) {
		// if it's not the value of a 0 (48.(rune)) set it
		col := i % 9
		row := (i - col) / 9
		b.Cells[row][col] = &cell{}
		if val > 48 && val < 58 {
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
	}

	b.rmNotPossibles()

	return nil
}

// spaceMap gradually increases the amount of allocated space as more
//	non-whitespace characters are encountered
func spaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// credit for the above spacemap function
// https://stackoverflow.com/a/32081891 CC BY-SA 4.0

func cellsAreValid(cells [9]*cell) bool {
	assertOnes := map[int]int{1: 0, 2: 0, 3: 0, 4: 0,
		5: 0, 6: 0, 7: 0, 8: 0, 9: 0}
	for _, cell := range cells {
		intVal := int(cell.Value)
		assertOnes[intVal]++
		if intVal != 0 && assertOnes[intVal] > 1 {
			fmt.Println("index:", intVal, "seems to be >1 occurrences")
			return false
		}
	}
	return true
}

// SolveHiddenSingles takes in a board and solves hidden singles until
//	it can find no more and returns the number of cells filled.
func SolveNakedSingles(b *board) int {
	solves := b.solveNakedSingles()
	sum := 0
	b.rmNotPossibles()
	for solves > 0 {
		solves = b.solveNakedSingles()
		sum += solves
		b.rmNotPossibles()
	}
	return sum
}

func SolveHiddenSingles(b *board) int {
	sum := 0
	solves := 0
	solves += b.solveHiddenSingles(getRow)
	b.rmNotPossibles()
	solves += b.solveHiddenSingles(getCol)
	b.rmNotPossibles()
	solves += b.solveHiddenSingles(getSet)
	b.rmNotPossibles()
	sum += solves
	for solves > 0 {
		solves = 0
		solves += b.solveHiddenSingles(getRow)
		b.rmNotPossibles()
		solves += b.solveHiddenSingles(getCol)
		b.rmNotPossibles()
		solves += b.solveHiddenSingles(getSet)
		b.rmNotPossibles()
		//fmt.Println("number of solves", solves)
		sum += solves
	}
	return sum
}
