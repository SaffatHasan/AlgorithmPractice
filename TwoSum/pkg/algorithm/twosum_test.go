package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected [2]int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: [2]int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: [2]int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: [2]int{0, 1}},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, TwoSum(tc.target, tc.nums))
	}
}
