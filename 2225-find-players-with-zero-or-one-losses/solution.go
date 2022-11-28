package findplayerswithzerooronelosses

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
	winners := map[int]int{}
	losers := map[int]int{}

	for _, m := range matches {
		winner := m[0]
		loser := m[1]

		winners[winner] = winners[winner] + 1
		losers[loser] = losers[loser] + 1
	}

	neverLost := []int{}
	for w := range winners {
		if _, isAlsoLoser := losers[w]; !isAlsoLoser {
			neverLost = append(neverLost, w)
		}
	}
	sort.Slice(neverLost, func(i, j int) bool {
		return neverLost[i] < neverLost[j]
	})

	lostOne := []int{}
	for l, count := range losers {
		if count == 1 {
			lostOne = append(lostOne, l)
		}
	}
	sort.Slice(lostOne, func(i, j int) bool {
		return lostOne[i] < lostOne[j]
	})

	return [][]int{neverLost, lostOne}
}
