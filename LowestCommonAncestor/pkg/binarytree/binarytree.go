package binarytree

const NIL_NODE int = -1

type Node struct {
	Val         int
	Left, Right *Node
}

// CreateTree generates a tree from an input array defined
// such that a node at index i has child elements at index 2*i and 2*i+1
// NOTE: nil children are represented by -1
func CreateTree(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	return createTreeHelper(arr, 0)
}

func createTreeHelper(arr []int, idx int) *Node {
	if idx >= len(arr) || arr[idx] == NIL_NODE {
		return nil
	}

	leftIdx := 2*idx + 1
	rightIdx := 2 * (idx + 1)

	return &Node{
		arr[idx],
		createTreeHelper(arr, leftIdx),
		createTreeHelper(arr, rightIdx),
	}
}
