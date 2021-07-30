package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleHistogram(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 1, 1},
			expected: 3,
		},
		{
			input:    []int{1},
			expected: 1,
		},
		{
			input:    []int{1, 2, 3},
			expected: 4,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: 6,
		},
		{
			input:    []int{1, 1},
			expected: 2,
		},
		{
			input:    []int{1, 1},
			expected: 2,
		},
		{
			input:    []int{1, 1, 3},
			expected: 3,
		},
		{
			input:    []int{2, 1, 3, 4, 5, 2, 6},
			expected: 10,
		},
		{
			input:    []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			input:    []int{1, 3, 3},
			expected: 6,
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, LargestRectangleHistogram(tc.input), "Input:%v", tc.input)
	}
}
