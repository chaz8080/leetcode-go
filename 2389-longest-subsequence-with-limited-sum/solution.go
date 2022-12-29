package longestsubsequencewithlimitedsum

import "sort"

func answerQueries(nums []int, queries []int) []int {
	res := make([]int, len(queries))
	sort.Ints(nums)

	for qIndex, query := range queries {
		sum := 0
		for _, num := range nums {
			sum += num
			res[qIndex]++
			if sum > query {
				res[qIndex]--
				break
			}
		}
	}

	return res
}
