package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubmatrixFull(t *testing.T) {
	testCases := []struct {
		input    [3][]int
		expected []bool
	}{
		{
			input: [3][]int{
				{1, 4, 7, 1},
				{2, 5, 8, 2},
				{3, 6, 9, 3},
			},
			expected: []bool{true, true},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, IsSubmatrixFull(tc.input))
	}
}

func TestIsAllUnique(t *testing.T) {
	num := 0
	for i := 0; i < 9; i++ {
		setBit(&num, i)
	}

	// in binary 0000000011111111
	// buckets 0->8 are set to 1
	assert.Equal(t, 511, num)
}
