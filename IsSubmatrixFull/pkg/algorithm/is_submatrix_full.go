package algorithm

func IsSubmatrixFull(input [3][]int) []bool {
	result := make([]bool, len(input[0])-2)
	bucket := 0

	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		for colIndex := 0; colIndex < 3; colIndex++ {
			// set ith bit to true if input[i][j] = i
			// e.g. set the 1st bit to true if
			// 		input[i][j] = 1
			setBit(&bucket, input[rowIndex][colIndex]-1)
		}
	}

	result[0] = isAllUnique(bucket)

	for colIndex := 3; colIndex < len(input[0]); colIndex++ {
		// remove column that is leaving
		for rowIndex := 0; rowIndex < 3; rowIndex++ {
			unsetBit(&bucket, input[rowIndex][colIndex-3])
		}

		// add column that is entering
		for rowIndex := 0; rowIndex < 3; rowIndex++ {
			setBit(&bucket, input[rowIndex][colIndex])
		}
		result[colIndex-2] = isAllUnique(bucket)
	}

	return result
}

// returns true if bits 0-8 are set to 1
// returns false otherwise
func isAllUnique(bucket int) bool {
	// 511 (base 10) = 1111111111 (base 2)
	// 8 bits set to 1 in sequence
	return bucket == 511
}

func setBit(a *int, pos int) {
	*a = *a | (1 << pos)
}

func unsetBit(a *int, pos int) {
	*a = *a & (^(1 << pos))
}
