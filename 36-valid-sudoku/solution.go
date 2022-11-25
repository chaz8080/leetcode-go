package validSudoku

// dot is byte representation of '.'
const dot = 46

func isValidSudoku(board [][]byte) bool {
	subgroups := make([]map[byte]bool, 9)
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)

	for i := range board {
		subgroups[i] = make(map[byte]bool)
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
	}

	for rowIndex := range board {
		for colIndex := range board[0] {
			char := board[rowIndex][colIndex]
			if char == dot {
				continue
			}

			if _, seenBefore := rows[rowIndex][char]; seenBefore {
				return false
			}
			rows[rowIndex][char] = true

			if _, seenBefore := cols[colIndex][char]; seenBefore {
				return false
			}
			cols[colIndex][char] = true

			subgroupIndex := (rowIndex/3)*3 + colIndex/3
			if _, seenBefore := subgroups[subgroupIndex][char]; seenBefore {
				return false
			}
			subgroups[subgroupIndex][char] = true
		}
	}
	return true
}
