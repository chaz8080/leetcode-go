package minnumberofarrowstoburstballoons

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] || (points[i][0] == points[j][0] && points[i][1] < points[j][1])
	})
	arrowsCount := 0
	for index := 0; index < len(points); index++ {
		arrowsCount++
		end := points[index][1]
		for index < len(points)-1 && points[index+1][0] <= end {
			end = min(end, points[index+1][1])
			index++
		}
	}

	return arrowsCount
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
[[1 2] [3 4] [5 6] [7 8]]

**/
