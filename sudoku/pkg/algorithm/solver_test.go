package algorithm

import (
	"testing"

	"github.com/SaffatHasan/AlgorithmPractice/sudoku/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	originalBoardString := "864371259325849761971265843436192587198657432257483916689734125713528694542916378"
	originalBoard := model.FromString(originalBoardString)

	b := model.FromString(originalBoardString)
	assert.True(t, b.IsValid())

	result := Solve(b)
	assert.True(t, result)
	assert.Equal(t, b, originalBoard)

	for i := 1; i < 9; i++ {
		b.SetVal(i, 0, 0)
		result = Solve(b)
		assert.True(t, result)
		assert.Equal(t, b, originalBoard)
	}
}

func TestSolverWithUnsolvable(t *testing.T) {
	b := model.FromString("999999999999999999999999999999999999999999999999999999999999999999999999999999999")
	assert.False(t, Solve(b))
}

func TestSolvePuzzle(t *testing.T) {
	// this puzzle has a unique solution
	b := model.FromString("000260701680070090190004500820100040004602900050003028009300074040050036703018000")
	assert.True(t, Solve(b), b.String())

	solved := model.FromString("435269781682571493197834562826195347374682915951743628519326874248957136763418259")
	assert.Equal(t, solved, b)

}

func TestFindZero(t *testing.T) {
	b := model.FromString("099999999999999999999999999999999999999999999999999999999999999999999999999999999")
	row, col := findZero(b)
	assert.Equal(t, 0, row)
	assert.Equal(t, 0, col)

	b = model.FromString("999999999999999999999999999999999999999999999999999999999999999999999999999999999")
	row, col = findZero(b)
	assert.Equal(t, -1, row)
	assert.Equal(t, -1, col)
}
