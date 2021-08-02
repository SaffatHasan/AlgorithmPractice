package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	fizz, buzz, fizzbuzz := "Fizz", "Buzz", "FizzBuzz"
	testCases := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, fizz},
		{4, "4"},
		{5, buzz},
		{6, fizz},
		{7, "7"},
		{8, "8"},
		{9, fizz},
		{10, buzz},
		{12, fizz},
		{15, fizzbuzz},
		{18, fizz},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, FizzBuzz(tc.input), tc.expected)
		})
	}
}
