package algorithms

import (
  "container/heap"
  "github.com/SaffatHasan/AlgorithmPractice/KthOrderStatistic/pkg/datastructures"
)

// given an input array and a k
// return the kth smallest element
// assumptions:
//   1 <= k <= arr.length
//
// note that this algorithm is O(k*log(n))
// but approaches O(n*log(n)) as k approaches n
//
// this could be optimized to find the n-k largest
// value instead of the kth smallest if k > n/2
// AKA nth order statistic
//
// ref http://en.wikipedia.org/wiki/Order_statistic#Notation_and_examples
func KthOrderStatistic(arr []int, k int) int {
  h := &datastructures.MaxHeap{}

  for _, val := range arr {
    heap.Push(h, val)
    if h.Len() > k {
      heap.Pop(h)
    }
  }

  return heap.Pop(h).(int)
}