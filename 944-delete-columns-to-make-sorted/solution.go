package deletecolumnstomakesorted

func minDeletionSize(strs []string) int {
	cols := len(strs[0])
	toDelete := 0

	for col := 0; col < cols; col++ {
		for row := 1; row < len(strs); row++ {
			// prevChar := strs[row-1][col]
			// char := strs[row][col]
			if strs[row-1][col] > strs[row][col] {
				toDelete++
				break
			}
		}
	}

	return toDelete
}
