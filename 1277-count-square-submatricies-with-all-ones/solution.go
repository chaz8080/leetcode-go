package countsquaresubmatricieswithallones

func countSquares(matrix [][]int) int {
	var sum int

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] != 1 {
				continue
			}
			sum += maxSize(row, col, matrix, 0)
		}
	}

	return sum
}

func maxSize(row, col int, matrix [][]int, size int) int {
	if row+size > len(matrix)-1 {
		return size
	}

	if col+size > len(matrix[0])-1 {
		return size
	}

	for r := row; r <= row+size; r++ {
		for c := col; c <= col+size; c++ {
			if matrix[r][c] != 1 {
				return size
			}
		}
	}

	size++
	return maxSize(row, col, matrix, size)
}
