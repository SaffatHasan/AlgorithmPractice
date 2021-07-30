package algorithm

import (
	"github.com/SaffatHasan/LargestRectangleHistogram/pkg/model"
)

func LargestRectangleHistogram(arr []int) int {
	result := 0

	s := model.Stack{}
	for i := 0; i <= len(arr); i++ {
		cur := 0
		if i < len(arr) {
			cur = arr[i]
		}

		for s.Len() > 0 && arr[s.Peek()] >= cur {
			right := i
			heightIndex, _ := s.Pop()
			height := arr[heightIndex]

			left := 0
			if s.Len() > 0 {
				left = s.Peek() + 1
			}

			area := height * (right - left)
			if area > result {
				result = area
			}
		}
		s.Push(i)
	}
	return result
}
