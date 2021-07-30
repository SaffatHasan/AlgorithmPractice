package algorithm

// given two strings, return the longest common subsequence
func LongestCommonSubsequence(word1, word2 string) int {
	// memo[i][j] is the longest common subsequence
	// for both strings up to:
	//	index i in word1
	// 	index j in word2
	M := make([][]int, len(word1)+1)

	for i := 0; i <= len(word1); i++ {
		M[i] = make([]int, len(word2)+1)

		// base case
		M[i][0] = 0
	}

	for j := 0; j <= len(word2); j++ {
		M[0][j] = 0
	}

	// fill out table
	globalMax := 0

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			word1Char := word1[i-1]
			word2Char := word2[j-1]

			if word1Char == word2Char {
				M[i][j] = M[i-1][j-1] + 1
			} else {
				M[i][j] = Max(M[i-1][j], M[i][j-1])
			}

			globalMax = Max(globalMax, M[i][j])
		}
	}

	return globalMax
}

// given two integers, return the larger value
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
