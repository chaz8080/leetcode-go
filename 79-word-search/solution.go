package solution

import "fmt"

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == word[0] {
				// pop first letter
				werd := word[1:]
				if checkWord(board, werd, nil, row, col) {
					return true
				}
			}
		}
	}

	return false
}

func getKey(row, col int) string {
	return fmt.Sprintf("%d%d", row, col)
}

func checkWord(board [][]byte, word string, visited map[string]bool, row int, col int) bool {
	if visited == nil {
		// start traversing to see if solution
		visited = map[string]bool{}
		visited[getKey(row, col)] = true
	}

	fmt.Println(word)
	if len(word) == 0 {
		return true
	}

	char, word := word[0], word[1:]

	rows := len(board)
	cols := len(board[0])

	dirs := [][]int{
		{0, -1},
		{0, 1},
		{-1, 0},
		{1, 0},
	}

	for _, dir := range dirs {
		x, y := dir[0], dir[1]
		posX := row + x
		posY := col + y

		if _, previouslyVisited := visited[getKey(posX, posY)]; previouslyVisited {
			continue
		}

		// if in bounds
		if (posX) >= 0 && (posX) < rows && (posY) >= 0 && (posY) < cols {
			if board[posX][posY] == char {
				visited[getKey(posX, posY)] = true
				if checkWord(board, word, visited, posX, posY) {
					return true
				}
			}
		}
	}

	return false
}
