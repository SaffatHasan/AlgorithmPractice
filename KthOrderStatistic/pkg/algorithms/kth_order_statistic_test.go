package algorithms

import (
  "testing" 
  "github.com/stretchr/testify/assert"
)

func TestKthOrderStatistic(t *testing.T) {
  testCases := []struct{
    input []int
    k int
    expected int
    description string
  }{
    {
      input: []int{1,2,3},
      k: 3,
      expected: 3,
      description: "3rd order on sorted array",
    },
    {
      input: []int{1,2,3},
      k: 2,
      expected: 2,
      description: "2nd order on sorted array",
    },
    {
      input: []int{1,2,3},
      k: 1,
      expected: 1,
      description: "1st order on sorted array",
    },
    {
      input: []int{1,5,3,2,4},
      k: 1,
      expected: 1,
      description: "1st order on unsorted array",
    },
    {
      input: []int{1,5,3,2,4},
      k: 2,
      expected: 2,
      description: "1st order on unsorted array",
    },
    {
      input: []int{1,5,3,2,4},
      k: 5,
      expected: 5,
      description: "1st order on unsorted array",
    },
  }

  for _, tc := range testCases {
    assert.Equal(
      t,
      tc.expected,
      KthOrderStatistic(tc.input, tc.k),
      "Test Case: %s failed",
      tc.description,
    )
  }
}