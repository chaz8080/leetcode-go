package twosumclosestk

import (
	"math"
	"sort"
)

func twosumclosestk(nums []int, target int) [2]int {
	res := [2]int{-1, -1}

	// assuming they're not already sorted, sort the numbers smallest to largest
	sort.Ints(nums)

	distanceFromTarget := math.MaxInt

	left := 0
	right := len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			res[0] = left + 1
			res[1] = right + 1
			return res
		}

		if abs(target-sum) < distanceFromTarget {
			res[0] = left + 1
			res[1] = right + 1
			distanceFromTarget = abs(target - sum)
		}

		if sum > 0 {
			right--
		} else {
			left++
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
