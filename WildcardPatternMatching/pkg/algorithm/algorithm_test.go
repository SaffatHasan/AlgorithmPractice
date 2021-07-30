package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	testCases := []struct {
		input    string
		pattern  string
		expected bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b", true},
		{"acdcb", "a*c?b", false},
		{"foo", "f*o*", true},
		{"fodo", "f*o*", true},
		{"boo", "f*o*", false},
		{"t", "f*o*", false},
		{"harry potter", "harry*", true},
		{"myemail@gmail.com", "*@*.com", true},
	}
	for _, tc := range testCases {
		actual := IsMatch(tc.input, tc.pattern)
		assert.Equal(t, actual, tc.expected, "IsMatch(input=%s, pattern=%s)=%t but expected %t", tc.input, tc.pattern, actual, tc.expected)
	}
}
