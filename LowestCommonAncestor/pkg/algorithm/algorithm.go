package algorithm

import "github.com/SaffatHasan/AlgorithmPractice/LowestCommonAncestor/pkg/binarytree"

// LowestCommonAncestor returns the lowest node which has `p` and `q` in tree `T`
// as its descendants where we define a node to be its own descendant.
func LowestCommonAncestor(tree *binarytree.Node, p, q int) *binarytree.Node {
	// base case
	if tree == nil {
		return nil
	}

	// from child
	left := LowestCommonAncestor(tree.Left, p, q)
	right := LowestCommonAncestor(tree.Right, p, q)

	// to parent
	// case 1: both left & right have found the node -> we are the LCA
	if left != nil && right != nil {
		return tree
	}

	// case 2: we match either p or q, propagate ourselves up
	if tree.Val == p || tree.Val == q {
		return tree
	}

	// case 3: return non-nil found child
	if right != nil {
		return right
	}

	if left != nil {
		return left
	}

	// case 4: return nil if nothing found
	return nil
}
