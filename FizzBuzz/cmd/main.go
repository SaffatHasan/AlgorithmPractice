package main

import (
	"fmt"

	"github.com/SaffatHasan/AlgorithmPractice/FizzBuzz/algorithm"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("FizzBuzz(%2d) -> %4s\n", i, algorithm.FizzBuzz(i))
	}
}
