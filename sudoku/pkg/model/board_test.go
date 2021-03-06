package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	row := genSequentialArr()
	board := genBoard(row)
	for idx := 0; idx < 9; idx++ {
		assert.Equal(t, row, board.GetRow(idx))
	}
}

func TestGetCol(t *testing.T) {
	row := genSequentialArr()
	board := genBoard(row)

	for idx := 0; idx < 9; idx++ {
		assert.Equal(t, genArrUsingConst(idx+1), board.GetCol(idx))
	}
}

func TestGetBox(t *testing.T) {
	board := genBoard(genSequentialArr())

	for idx := 0; idx < 9; idx++ {
		assert.Equal(t, genBox(idx), board.getBox(idx), "BoxIndex = %d\n%v\n", idx, board.String())
	}
}

func genBoard(row [9]int) Board {
	return Board{
		[9][9]int{
			row, row, row,
			row, row, row,
			row, row, row,
		},
	}

}

func genSequentialArr() [9]int {
	return [9]int{
		1, 2, 3,
		4, 5, 6,
		7, 8, 9,
	}
}

func genArrUsingConst(x int) [9]int {
	return [9]int{
		x, x, x,
		x, x, x,
		x, x, x,
	}
}

func genBox(boxIndex int) [9]int {
	offset := boxIndex % 3
	return [9]int{
		1 + 3*offset, 2 + 3*offset, 3 + 3*offset,
		1 + 3*offset, 2 + 3*offset, 3 + 3*offset,
		1 + 3*offset, 2 + 3*offset, 3 + 3*offset,
	}
}

func TestIsValid(t *testing.T) {
	b := genBoard(genArrUsingConst(0))
	assert.True(t, b.IsValid())

	b = genBoard(genSequentialArr())
	assert.False(t, b.IsValid())
}

func TestIsValidHelper(t *testing.T) {
	testcases := []struct {
		input    [9]int
		expected bool
	}{
		{[9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, true},
		{[9]int{9, 2, 3, 4, 5, 6, 7, 8, 9}, false},
		{[9]int{1, 3, 3, 4, 5, 6, 7, 8, 9}, false},
	}

	for _, tc := range testcases {
		assert.Equal(t, tc.expected, isValid(tc.input))
	}
}
