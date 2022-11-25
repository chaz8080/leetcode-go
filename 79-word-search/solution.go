package solution

type Data struct {
	Board [][]byte
	Rows  int
	Cols  int
}

func exist(board [][]byte, word string) bool {
	data := &Data{
		Board: board,
		Rows:  len(board),
		Cols:  len(board[0]),
	}

	for row := 0; row < data.Rows; row++ {
		for col := 0; col < data.Cols; col++ {
			if board[row][col] == word[0] {
				if checkWord(data, row, col, word) {
					return true
				}
			}
		}
	}

	return false
}

func checkWord(data *Data, row int, col int, word string) bool {
	// base case: when we've found all of the letters in the word
	if len(word) == 0 {
		return true
	}

	dirs := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	if row < 0 || row == data.Rows || col < 0 || col == data.Cols || data.Board[row][col] != word[0] {
		return false
	}

	ret := false
	data.Board[row][col] = '#'
	for _, dir := range dirs {
		x, y := dir[0], dir[1]
		posX := row + x
		posY := col + y
		ret = checkWord(data, posX, posY, word[1:])

		if ret {
			break
		}
	}

	data.Board[row][col] = word[0]

	return ret
}
