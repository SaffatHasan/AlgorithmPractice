package main

import (
	"fmt"

	"github.com/SaffatHasan/AlgorithmPractice/LargestRectangleHistogram/pkg/algorithm"
)

func main() {
	input := []int{1, 2, 3}
	fmt.Printf("The largest histogram found in the histogram:\n%v\nhas area: %d.", input, algorithm.LargestRectangleHistogram(input))
}
