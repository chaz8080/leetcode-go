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
	if row+size > len(matrix)-1 || col+size > len(matrix[0])-1 {
		return size
	}

	// check all row, last col
	for r := row; r < row+size; r++ {
		if matrix[r][col+size] != 1 {
			return size
		}
	}

	// check last row
	for c := col; c <= col+size; c++ {
		if matrix[row+size][c] != 1 {
			return size
		}
	}

	size++
	return maxSize(row, col, matrix, size)
}
