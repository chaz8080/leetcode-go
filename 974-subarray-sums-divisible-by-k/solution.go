package subarraysumsdivisiblebyk

import "log"

func subarraysDivByK(nums []int, k int) int {
	if len(nums) == 1 {
		if nums[0]%k == 0 {
			return 1
		} else {
			return 0
		}
	}

	totsCount := 0
	contCount := 0
	for i := len(nums) - 1; i > 0; i-- {
		sums := make([]int, len(nums)-i)

		for j := len(nums) - len(sums); j < len(nums); j++ {
			log.Println(sums)
			if nums[j]%k == 0 {
				contCount++
				continue
			}

			if contCount > 0 {
				
			}
		}
	}

	return totsCount
}
