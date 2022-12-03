package determineIfStringHalvesAreClose

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	mW1 := getMapSymbolsNumbers(word1)
	mW2 := getMapSymbolsNumbers(word2)

	if !compareSymbolsContains(mW1, mW2) {
		return false
	}

	mNumW1 := getNumbersMap(mW1)
	mNumW2 := getNumbersMap(mW2)

	return compareNumbersContains(mNumW1, mNumW2)

}

func getNumbersMap(w map[byte]int) map[int]int {
	m := make(map[int]int)
	for _, val := range w {
		m[val] += 1
	}
	return m
}

func compareNumbersContains(w1, w2 map[int]int) bool {
	if len(w1) != len(w2) {
		return false
	}
	for key, val := range w1 {
		if _, ok := w2[key]; !ok || w2[key] != val {
			return false
		}
	}
	return true
}

func compareSymbolsContains(w1, w2 map[byte]int) bool {
	if len(w1) != len(w2) {
		return false
	}
	for key := range w1 {
		if _, ok := w2[key]; !ok {
			return false
		}
	}
	return true
}

func getMapSymbolsNumbers(w string) map[byte]int {
	m := make(map[byte]int)
	for i := range w {
		m[w[i]] += 1
	}
	return m
}
