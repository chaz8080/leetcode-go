package maxpointsonaline

import (
	"math"
)

func maxPoints(points [][]int) int {
	lp := len(points)
	if lp == 1 {
		return 1
	}
	res := 2

	for i := 0; i < lp; i++ {
		counts := make(map[float64]int)
		for j := 0; j < lp; j++ {
			if i == j {
				continue
			}
			key := math.Atan2(float64(points[j][1])-float64(points[i][1]), float64(points[j][0])-float64(points[i][0]))
			counts[key]++
		}

		for _, count := range counts {
			if res < count+1 {
				res = count + 1
			}
		}
	}

	return res
}
