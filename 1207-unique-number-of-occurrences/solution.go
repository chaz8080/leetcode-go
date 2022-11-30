package uniquenumberofoccurrences

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}

	for _, v := range arr {
		m[v] = m[v] + 1
	}

	s := map[int]bool{}
	for _, v := range m {
		s[v] = true
	}

	return len(m) == len(s)
}
