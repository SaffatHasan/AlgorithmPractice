package algorithm

func IsMatch(input, pattern string) bool {
	// Human Language Explanation
	//
	// Start with the base case of matching the empty string against an empty pattern
	// Next, expand by
	//		Case 1: expand by matching a pattern literal or '?' to an input literal (consume one additional pattern and input character)
	//		Case 2: expand by matching '*' to no characters 						(consume one additional pattern character, the '*')
	//		Case 3: expand by matching '*' to one additional character 				(consume one additional input 	character, but no pattern characters)

	// Dynamic Programming - Finite State Machine Explanation
	//
	// state definition:
	// 		i (int): index in _input_ (exclusive) up to which the current state matches
	//		j (int): index in _pattern_ (exclusive) up to which the current state matches
	//		is_match (bool): whether the current substring of input and pattern match
	//
	// 		e.g. input="foo", pattern="f*o*" i=1, j=1
	//		input[:i]="f", pattern[:j]="f" is true as the f literal matches f

	// state transition
	//		literal character matches -> M[i][j] = M[i-1][j-1]
	//		'?'		character matches -> M[i][j] = M[i-1][j-1]
	//		'*' matches no characters -> M[i][j] = M[i][j-1]
	//		'*' matches 1+ characters -> M[i][j] = M[i-1][j]

	// base case
	//		M[0][0] = true  -> empty pattern matches empty string
	// 		M[i][0]	= false -> empty pattern cannot match non-empty string
	//		M[0][j] = true if pattern[j-1] == '*' and M[0][j-1] is true
	//   		      else false

	// setup memoization table
	M := make([][]bool, len(input)+1)

	for i := range M {
		M[i] = make([]bool, len(pattern)+1)
	}

	// base case
	M[0][0] = true

	for i := 1; i <= len(input); i++ {
		M[i][0] = false
	}

	for j := 1; j <= len(pattern); j++ {
		if pattern[j-1] == '*' {
			M[0][j] = M[0][j-1]
		} else {
			M[0][j] = false
		}
	}

	// populate memoization table
	for i := 1; i <= len(input); i++ {
		for j := 1; j <= len(pattern); j++ {
			if pattern[j-1] == input[i-1] || pattern[j-1] == '?' {
				M[i][j] = M[i-1][j-1]
			} else if pattern[j-1] == '*' {
				M[i][j] = M[i-1][j] || M[i][j-1]
			}
		}
	}
	return M[len(input)][len(pattern)]
}
