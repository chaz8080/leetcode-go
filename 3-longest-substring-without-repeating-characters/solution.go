package longestsubstringwithoutrepeatingcharacters

// notes: sliding window with a hashmap

func lengthOfLongestSubstring(s string) int {
	// keep a map of the characters seen to the index they occur
	charIndex := make(map[byte]int)

	// keep track of max
	maxLen := 0

	// create two pointers to keep track of where we are in the array
	var left, right int

	// until the right pointer gets to the end, keep iterating
	for right < len(s) {
		char := s[right]

		index, seen := charIndex[char]
		if seen && index >= left {
			// move the left pointer to the right of the repeated char
			left = index + 1
		}

		charIndex[char] = right
		right++
		maxLen = max(maxLen, right-left)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
