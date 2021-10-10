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
	Digit     uint8
	DigitSet  bool
	Candidate map[uint8]bool
}

func (c *cell) setEmpty() {
	c.Digit = 0
	c.DigitSet = false
	c.Candidate = map[uint8]bool{
		1: true, 2: true, 3: true, 4: true, 5: true,
		6: true, 7: true, 8: true, 9: true}
}

func (c *cell) set(value uint8) {
	c.Digit = value
	c.DigitSet = true
	c.Candidate = map[uint8]bool{
		1: false, 2: false, 3: false, 4: false, 5: false,
		6: false, 7: false, 8: false, 9: false}
	c.Candidate[value] = true
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
			if cell != nil && cell.Digit != 0 {
				fmt.Printf(pad+"%d"+pad, cell.Digit)
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

func (b *board) rmNotCandidates() {
	// loop over each row and col
	//	and remove candidates in row, col, set for set values
	for rowIdx := 0; rowIdx < 9; rowIdx++ {
		for colIdx := 0; colIdx < 9; colIdx++ {
			// only need to set not candidates for the cells with set values
			if b.Cells[rowIdx][colIdx].DigitSet {
				setIdx := computeSetIdx(rowIdx, colIdx)
				val := b.Cells[rowIdx][colIdx].Digit
				b.rmNotCandidatesSet(colIdx, val, getCol)
				b.rmNotCandidatesSet(rowIdx, val, getRow)
				b.rmNotCandidatesSet(setIdx, val, getSet)
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
			if !myCell.DigitSet {
				// count the number of candidate values, keeping the last one
				candidateCount := 0
				lastCandidate := uint8(0)
				for value, candidate := range myCell.Candidate {
					if candidate {
						candidateCount++
						lastCandidate = value
					}
				}
				// if there's only one set the last value for the cell
				if candidateCount == 1 {
					myCell.set(uint8(lastCandidate))
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
	// for each (row|col|set) get all the cells
	for idx := 0; idx < 9; idx++ {
		cells := getter(b, idx)
		var digit uint8
		// for each candidate value of a cell, check all the cells
		for digit = 1; digit < 10; digit++ {
			for cIdx := 0; cIdx < 9; cIdx++ {
				candidate := cells[cIdx]
				if candidate.DigitSet {
					continue
				}
				numCandidates := 0
				// for each other cell check our cell is not the one
				//	being compared against, if there's no other candidates
				//	our cell should be set
				for otherIdx := 0; otherIdx < 9; otherIdx++ {
					other := cells[otherIdx]
					if candidate == other {
						continue
					}
					if other.Candidate[digit] {
						numCandidates++
					}
				}
				if numCandidates == 0 {
					candidate.set(digit)
					solvedCells++
					break
				}
			}
		}
	}
	return solvedCells
}

func (b *board) rmNotCandidatesSet(idx int, value uint8,
	getter func(*board, int) [9]*cell) {
	set := getter(b, idx)
	for i := 0; i < 9; i++ {
		if !set[i].DigitSet {
			set[i].Candidate[value] = false
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
			b.Cells[row][col].set(parsedVal)
		} else if val == 48 {
			b.Cells[row][col].setEmpty()
		} else {
			return errors.New("expected rune 48-57; got: " +
				strconv.Itoa(int(val)))
		}
	}

	b.rmNotCandidates()

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
		intVal := int(cell.Digit)
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
	b.rmNotCandidates()
	for solves > 0 {
		solves = b.solveNakedSingles()
		sum += solves
		b.rmNotCandidates()
	}
	return sum
}

func SolveHiddenSingles(b *board) int {
	sum := 0
	solves := 1
	getters := []func(*board, int) [9]*cell{getRow, getCol, getSet}
	for solves > 0 {
		solves = 0
		for _, getter := range getters {
			solves += b.solveHiddenSingles(getter)
			b.rmNotCandidates()
		}
		sum += solves
	}
	return sum
}
