package maxpointsonaline

import (
	"fmt"
	"log"
	"math"
)

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	max := 0
	counts := make(map[string]int)
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			m := slope(points[i][0], points[i][1], points[j][0], points[j][1])
			yi := yint(points[i][0], points[i][1], m)
			key := fmt.Sprintf("%d%d", m, yi)
			if _, exists := counts[key]; !exists {

				counts[key] += 2
			} else {
				counts[key]++
			}

			if counts[key] > max {
				max = counts[key]
			}
			log.Println(counts)
		}
	}

	if max > len(points) {
		max--
	}

	return max
}

func slope(x1, y1, x2, y2 int) int {
	if x2-x1 == 0 {
		return math.MaxInt
	}
	decimalPlaces := 4
	return (y2 - y1) * 10 * decimalPlaces / (x2 - x1)
}

func yint(x, y, m int) int {
	decimalPlaces := 4
	return y - m/decimalPlaces*x
}
