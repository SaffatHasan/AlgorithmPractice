package model

const NIL_NODE int = -1

type Node struct {
	data        int
	left, right *Node
}

// CreateTree generates a tree from an input array defined
// such that a node at index i has child elements at index 2*i and 2*i+1
// NOTE: nil children are represented by -1
func CreateTree(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	return CreateTreeHelper(arr, 0)
}

func CreateTreeHelper(arr []int, idx int) *Node {
	if idx >= len(arr) || arr[idx] == NIL_NODE {
		return nil
	}

	leftIdx := 2*idx + 1
	rightIdx := 2 * (idx + 1)

	return &Node{
		arr[idx],
		CreateTreeHelper(arr, leftIdx),
		CreateTreeHelper(arr, rightIdx),
	}
}
