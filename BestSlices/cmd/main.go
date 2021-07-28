package main

import (
	"fmt"

	"github.com/SaffatHasan/BestSlices/pkg/algorithm"
)

func main() {
	pizzaSlices := []int{1, 2, 100, 4, 3, 5, 6}
	fmt.Printf("BestSlicesSum(%v)=%d", pizzaSlices, algorithm.BestSlicesSum(pizzaSlices))

}
