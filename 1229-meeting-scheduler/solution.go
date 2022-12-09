package meetingscheduler

import "sort"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	res := []int{}

	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})

	// use 2 pointers p1 and p2 respectively.
	p1, p2 := 0, 0
	for p1 < len(slots1) && p2 < len(slots2) {
		s1 := slots1[p1]
		s2 := slots2[p2]

		if s1[1]-s1[0] < duration {
			p1++
			continue
		}
		if s2[1]-s2[0] < duration {
			p2++
			continue
		}
		if s1[1] < s2[0] {
			p1++
			continue
		}
		if s2[1] < s1[0] {
			p2++
			continue
		}
		if min(s1[1], s2[1])-max(s1[0], s2[0]) >= duration {
			start := max(s1[0], s2[0])
			end := start + duration
			return []int{start, end}
		} else {
			if max(s1[1], s2[1]) == s1[1] {
				p2++
			} else {
				p1++
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
