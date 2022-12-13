package threesum

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}

	sort.Ints(nums)

	twoSum := func(nums []int, index int) {
		low := index + 1
		high := len(nums) - 1
		for low < high {
			sum := nums[index] + nums[low] + nums[high]
			if sum < 0 {
				low++
			} else if sum > 0 {
				high--
			} else {
				res = append(res, []int{nums[index], nums[low], nums[high]})
				low++
				high--
				for low < high && nums[low] == nums[low-1] {
					low++
				}
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] != nums[i] {
			twoSum(nums, i)
		}
	}

	return res
}
