package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestEqualityOperatorWorksAsExpected(t *testing.T) {
// 	treeA := Node{
// 		left: &Node{nil, nil, 5},
// 	}

// 	treeB := Node{
// 		left: &Node{nil, nil, 5},
// 	}

// 	assert.Equal(t, treeA, treeB)

// 	treeA.right = &treeB

// 	assert.NotEqual(t, treeA, treeB)
// }

func TestCreateTree(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected *Node
	}{
		{"Empty array returns empty", []int{}, nil},
		{"Single element array returns one node", []int{1}, &Node{1, nil, nil}},
		{
			"Two elements populates left child",
			[]int{3, 4},
			&Node{
				3,
				&Node{4, nil, nil},
				nil,
			},
		},
		{
			"Fully populated 3 layer tree",
			[]int{1, 2, 3, 4, 5, 6, 7},
			&Node{
				1,
				&Node{
					2,
					&Node{4, nil, nil},
					&Node{5, nil, nil},
				},
				&Node{
					3,
					&Node{6, nil, nil},
					&Node{7, nil, nil},
				},
			},
		},
		{
			"3 layer tree with empty nodes",
			[]int{1, NIL_NODE, 3, NIL_NODE, NIL_NODE, 6},
			&Node{
				1,
				nil,
				&Node{
					3,
					&Node{6, nil, nil},
					nil,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := CreateTree(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
