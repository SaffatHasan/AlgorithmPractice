package main

import (
	"fmt"

	"github.com/SaffatHasan/LongestCommonSubsequence/pkg/algorithm"
)

func main() {
	word1, word2 := "foobar", "foobazbar"
	fmt.Printf(
		"LongestCommonSubsequence(\"%s\",\"%s\")=%d\n",
		word1,
		word2,
		algorithm.LongestCommonSubsequence(word1, word2),
	)
}
