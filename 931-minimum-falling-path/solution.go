package minimumfallingpath

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for row := 1; row < n; row++ {
		for col := 0; col < n; col++ {
			top := getVal(row, col, matrix) + getVal(row-1, col, matrix)
			diagonallyLeft := getVal(row, col, matrix) + getVal(row-1, col-1, matrix)
			diagonallyRight := getVal(row, col, matrix) + getVal(row-1, col+1, matrix)
			matrix[row][col] = min(top, diagonallyLeft, diagonallyRight)
		}
	}
	return min(matrix[n-1]...)
}

func getVal(row, col int, matrix [][]int) int {
	if (row >= 0 && row < len(matrix)) && (col >= 0 && col < len(matrix)) {
		return matrix[row][col]
	}
	return 10_000
}

func min(args ...int) int {
	ans := args[0]
	for _, v := range args {
		if ans > v {
			ans = v
		}
	}
	return ans
}
