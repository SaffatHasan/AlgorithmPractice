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
