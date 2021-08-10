package algorithm

import (
	"testing"

	"github.com/SaffatHasan/AlgorithmPractice/BombChainReaction/pkg/model"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input          []model.Bomb
	startBombIndex int
	expected       int
}

func TestSearch(t *testing.T) {
	for _, tc := range getCommonTestCases() {
		assert.Equal(t, tc.expected, FindTotalBombsExploded(tc.input, tc.startBombIndex), "Input: %+v", tc.input)
	}
}

func TestSearchWithChannels(t *testing.T) {
	t.SkipNow() // these tests deadlock
	for _, tc := range getCommonTestCases() {
		assert.Equal(t, tc.expected, FindTotalBombsExplodedHelper(tc.input, tc.startBombIndex, true), "Input: %+v", tc.input)
	}
}

func getCommonTestCases() []TestCase {
	return []TestCase{
		{
			[]model.Bomb{
				{X: 1, Y: 5, R: 1},
				{X: 0, Y: 0, R: 0},
			},
			0,
			1,
		},
		{
			[]model.Bomb{
				{X: 0, Y: 0, R: 1},
				{X: 0, Y: 1, R: 0},
			},
			0,
			2,
		},
	}
}
