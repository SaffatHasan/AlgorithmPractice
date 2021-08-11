package algorithm

// tc: O(n)
// sc: O(n)
func TwoSum(target int, nums []int) [2]int {
	// store complementary values in map
	// map[VALUE] = INDEX
	// where nums[INDEX] = target - value
	complementMap := make(map[int]int)

	for index, value := range nums {
		if complementIndex, found := complementMap[value]; found {
			// for current value we have an index for its complement
			return [2]int{complementIndex, index}
		} else {
			// calculate complement value
			complement := target - value
			complementMap[complement] = index
		}
	}
	// no result found
	return [2]int{}
}
