package twosumii

import (
	"sort"
)

func twosumii(nums []int, target int) [2]int {
	res := [2]int{-1, -1}

	// sort the numbers smallest to largest
	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			res[0] = left + 1
			res[1] = right + 1
			return res
		}

		if sum > target {
			right--
		} else {
			left++
		}
	}

	return res
}
