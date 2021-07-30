package main

import (
  "github.com/SaffatHasan/KthOrderStatistic/pkg/algorithms"
  "fmt"
)
func main() {
	arr := []int{9,1,2,6,3,5}
  k := 4

  fmt.Printf(
    "The %dth order statistic of %v is %d\n",
    k,
    arr,
    algorithms.KthOrderStatistic(arr, k),
  )
}