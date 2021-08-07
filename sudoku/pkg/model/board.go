package model

import "fmt"

type Board struct {
	data [9][9]int
}

// returns the 0-index row
func (b *Board) getRow(rowIndex int) [9]int {
	return b.data[rowIndex]
}

// returns the 0-index column
func (b *Board) getCol(colIndex int) (result [9]int) {
	for rowIndex := 0; rowIndex < 9; rowIndex++ {
		result[rowIndex] = b.data[rowIndex][colIndex]
	}
	return result
}

// returns the 0-index box as following this
// 0 1 2
// 3 4 5
// 6 7 8
func (b *Board) getBox(boxIndex int) (result [9]int) {
	counter := 0
	fmt.Printf("boxIndex: %v\n", boxIndex)
	startRow := 3 * (boxIndex / 3)
	startCol := 3 * (boxIndex % 3)
	for rowIdx := startRow; rowIdx < startRow+3; rowIdx++ {
		for colIdx := startCol; colIdx < startCol+3; colIdx++ {
			fmt.Printf("(%v, %v)\n", rowIdx, colIdx)
			result[counter] = b.data[rowIdx][colIdx]
			counter++
		}
	}
	return
}

func (b *Board) isValid() bool {
	for idx := 0; idx < 9; idx++ {
		if !isValid(b.getRow(idx)) {
			return false
		}
		if !isValid(b.getCol(idx)) {
			return false
		}
		if !isValid(b.getBox(idx)) {
			return false
		}
	}
	return true
}

// returns true if 1-9 are in the array
// return false if any are missing
func isValid(data [9]int) bool {
	// can be optimized using bit manipulation
	// see https://www.geeksforgeeks.org/find-duplicates-of-array-using-bit-array/
	seen := make(map[int]bool)

	for _, val := range data {
		// 0 is set as placeholder, ignore it
		if val == 0 {
			continue
		}
		if seen[val] {
			return false
		}
		seen[val] = true
	}
	return true
}

func (b *Board) String() (result string) {
	separator := "+-------+-------+-------+\n"
	rowString := "| %d %d %d | %d %d %d | %d %d %d |\n"
	for rowIndex := 0; rowIndex < 9; rowIndex++ {
		if rowIndex%3 == 0 {
			result += separator
		}
		result += formatRow(rowString, b.getRow(rowIndex))
	}
	result += separator
	return
}

func formatRow(rowString string, row [9]int) string {
	return fmt.Sprintf(rowString, row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8])
}
