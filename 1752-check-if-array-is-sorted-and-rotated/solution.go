package checkIfArrayIsSortedAndRotated

func check(nums []int) bool {
	var min int = 1e2
	minIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			minIndex = i
		}
	}

	inc := 0
	for j := minIndex; j < len(nums)-1; j++ {
		if nums[j] > nums[j+1] {
			inc++
		}

		if inc > 1 {
			return false
		}
	}

	if nums[len(nums)-1] > nums[0] {
		inc++

		if inc > 1 {
			return false
		}
	}

	for k := 0; k < minIndex; k++ {
		if nums[k] > nums[k+1] {
			inc++
		}

		if inc > 1 {
			return false
		}
	}

	return true
}
