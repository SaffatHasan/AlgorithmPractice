package main

import (
	"fmt"

	"github.com/SaffatHasan/AlgorithmPractice/WildcardPatternMatching/pkg/algorithm"
)

func main() {
	input, pattern := "foo", "f*o*"
	fmt.Printf(
		"isMatch(input='%s', pattern='%s')=%t\n",
		input,
		pattern,
		algorithm.IsMatch(input, pattern),
	)
}
