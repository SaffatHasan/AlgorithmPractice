package algorithm

import "github.com/SaffatHasan/AlgorithmPractice/sudoku/pkg/model"

func Solve(b *model.Board) bool {
	// base case 1: board is not valid, return false
	if !b.IsValid() {
		return false
	}

	zeroRow, zeroCol := findZero(b)

	// base case 2: board is valid and no zeros -> done
	if zeroCol == -1 {
		return true
	}

	for i := 1; i <= 9; i++ {
		// set value
		b.SetVal(zeroRow, zeroCol, i)
		if Solve(b) {
			return true
		}
	}

	// unset and backtrack
	b.SetVal(zeroCol, zeroRow, 0)
	return false
}

func findZero(b *model.Board) (int, int) {
	zeroRow, zeroCol := 0, 0

	for i := zeroRow; i < 9; i++ {
		row := b.GetRow(i)
		for j := zeroCol; j < 9; j++ {
			if row[j] == 0 {
				return i, j
			}
		}
	}

	return -1, -1
}
