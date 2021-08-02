package algorithm

import (
	"testing"

	. "github.com/SaffatHasan/AlgorithmPractice/LowestCommonAncestor/pkg/binarytree"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		tree     *Node
		p        int
		q        int
		expected int
	}{
		{
			desc: "Example 1",
			tree: CreateTree(
				[]int{3, 5, 1, 6, 2, 0, 8, NIL_NODE, NIL_NODE, 7, 4},
			),
			p:        5,
			q:        1,
			expected: 3,
		},
		{
			desc: "Example 2",
			tree: CreateTree(
				[]int{3, 5, 1, 6, 2, 0, 8, NIL_NODE, NIL_NODE, 7, 4},
			),
			p:        5,
			q:        4,
			expected: 5,
		},
		{
			desc: "Example 3",
			tree: CreateTree(
				[]int{1, 2},
			),
			p:        1,
			q:        2,
			expected: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actual := LowestCommonAncestor(tc.tree, tc.p, tc.q).Val
			assert.Equal(t, tc.expected, actual)
		})
	}
}
