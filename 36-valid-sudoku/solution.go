package solution

import "log"

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

			if _, found := rows[rowIndex][char]; found {
				return false
			} else {
				rows[rowIndex][char] = true
			}

			if _, found := cols[colIndex][char]; found {
				return false
			} else {
				cols[colIndex][char] = true
			}

			sg := getSubgroup(rowIndex, colIndex)
			if _, found := subgroups[sg][char]; found {
				log.Println(rowIndex, colIndex)
				log.Println(subgroups)
				return false
			} else {
				subgroups[sg][char] = true
			}
		}
	}
	return true
}

func getSubgroup(x, y int) int {
	if x <= 2 && y <= 2 {
		return 0
	}
	if x <= 2 && y >= 3 && y <= 5 {
		return 1
	}
	if x <= 2 && y >= 6 && y <= 8 {
		return 2
	}
	if x >= 3 && x <= 5 && y <= 2 {
		return 3
	}
	if x >= 3 && x <= 5 && y >= 3 && y <= 5 {
		return 4
	}
	if x >= 3 && x <= 5 && y >= 6 && y <= 8 {
		return 5
	}
	if x >= 6 && y <= 2 {
		return 6
	}
	if x >= 6 && y >= 3 && y <= 5 {
		return 7
	}
	return 8
}
