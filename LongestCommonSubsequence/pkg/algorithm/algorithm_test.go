package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	testCases := []struct {
		word1    string
		word2    string
		expected int
	}{
		{"abc", "abc", 3},
		{"justice league", "league of legends", 7},
		{"abcd", "ffabcd", 4},
		{"stuff", "ffuts", 2},
		{"saffathasan", "longlongmou", 1},
	}
	for _, tc := range testCases {
		actual := LongestCommonSubsequence(tc.word1, tc.word2)
		assert.Equal(
			t,
			actual,
			tc.expected,
			"Got %d but expected %d for word1='%s', word2='%s'",
			actual,
			tc.expected,
			tc.word1,
			tc.word2,
		)
	}
}
