package algorithm

// Determine the largest sum of pizza slices you can get
// given an array of pizza slices where you and a friend
// take turns selecting a slice.
// You can pick a slice at either end of the list but your
// friend always picks the largest of the two ends.
// Good luck!
func BestSlicesSum(input []int) int {
	// generate N x N matrix
	M := GenMatrix(len(input))

	// base case 1: 1 move, we go first
	for i := 0; i < len(input); i++ {
		M[i][i] = input[i]
	}

	// base case 2: 2 moves, we take the larger slice then opponent takes remaining slice
	for i := 0; i+1 < len(input); i++ {
		M[i][i+1] = Max(M[i][i], M[i+1][i+1])
	}

	// diagonal iteration
	// 4 cases: LL, LR, RL, RR
	// we choose the first L/R
	// the opponent's choice is determined by the ends of the array
	// i.e. i+1->j if we chose the left  (ith) slice
	// i.e. i->j+1 if we chose the right (jth) slice
	for offset := 2; offset < len(input); offset++ {
		for i := 0; i+offset < len(input); i++ {
			j := i + offset

			// compute left
			left := input[i]
			if input[i+1] > input[j] {
				left += M[i+2][j] // opponent chose input[i+1], we get input[i] and our bestSlices in range i+2->j
			} else {
				left += M[i+1][j-1] // opponent chose input[j], we get input[i] and our bestSlices in range i+1->j-1
			}

			// compute right
			right := input[j]
			if input[i] < input[j-1] {
				right += M[i][j-2]
			} else {
				right += M[i+1][j-1]
			}

			M[i][j] = Max(left, right)
		}
	}
	return M[0][len(input)-1]
}

// generate a 2D matrix of size n x n
// of integers
func GenMatrix(n int) [][]int {
	M := make([][]int, n)
	for idx := range M {
		M[idx] = make([]int, n)
	}
	return M
}

// why doesn't golang implement this?
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
