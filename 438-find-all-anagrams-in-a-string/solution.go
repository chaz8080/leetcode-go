package findallanagramsinastring

func findAnagrams(s string, p string) []int {
	var res []int

	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return []int{}
	}

	// store the number of occurrences of each character in p and s[0...pLen]
	var pCount, sCount [26]int
	for i := 0; i < pLen; i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	// iterate over all possible substrings of s
	for i := 0; i <= sLen-pLen; i++ {
		if pCount == sCount {
			// found an anagram
			res = append(res, i)
		}

		// use a sliding window
		sCount[s[i]-'a']--
		if i+pLen < sLen {
			sCount[s[i+pLen]-'a']++
		}
	}

	return res
}
