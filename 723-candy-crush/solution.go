package candycrush

import (
	"log"
	"math"
)

func candyCrush(board [][]int) [][]int {
	// flag to know whether we should process the board again
	bubblesBurst := false

	// for each row, use a sliding window, size 3, to find any triplicates
	for row := 0; row < len(board); row++ {
		for col := 2; col < len(board[row]); col++ {
			a, b, c := math.Abs(float64(board[row][col-2])), math.Abs(float64(board[row][col-1])), math.Abs(float64(board[row][col]))
			if a == 0 || b == 0 || c == 0 {
				continue
			}
			if a == b && b == c {
				bubblesBurst = true
				if board[row][col-2] > 0 {
					board[row][col-2] = -1 * board[row][col-2]
				}
				if board[row][col-1] > 0 {
					board[row][col-1] = -1 * board[row][col-1]
				}
				if board[row][col] > 0 {
					board[row][col] = -1 * board[row][col]
				}
			}
		}
	}

	// for each col, use a sliding window, size 3, to find any triplicates
	for row := 2; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			a, b, c := math.Abs(float64(board[row-2][col])), math.Abs(float64(board[row-1][col])), math.Abs(float64(board[row][col]))
			if a == 0 || b == 0 || c == 0 {
				continue
			}
			if a == b && b == c {
				bubblesBurst = true
				log.Println(board[row-2][col], board[row-1][col], board[row][col])
				if board[row-2][col] > 0 {
					board[row-2][col] = -1 * board[row-2][col]
				}
				if board[row-1][col] > 0 {
					board[row-1][col] = -1 * board[row-1][col]
				}
				if board[row][col] > 0 {
					board[row][col] = -1 * board[row][col]
				}
			}
		}
	}

	// zero out all negative numbers (bubbles burting)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] < 0 {
				board[row][col] = 0
			}
		}
	}

	/**
	handle gravity, 0's bubble to the top

	iterate from bottom row to the top and bubble the 0's up

	think about it like this:

	imagine this is a column: [2, 0, 3, 4, 5, 0, 3]
	* both pointer @ 2 (index 0)
	* move both until hitting a zero
		bottom stays and top keeps iterating
			if top == end
				if bottom != end, fill the rest 0's
			if top next == 0 keep iterating
			if top next != 0 copy val to bottom, and move bottom to next
	*/

	for col := 0; col < len(board[0]); col++ {
		top := len(board) - 1
		bottom := len(board) - 1
		for top >= 0 {
			if board[top][col] == 0 {
				top--
			} else {
				board[bottom][col] = board[top][col]
				top--
				bottom--
			}
		}

		// fill the remaining
		for bottom >= 0 {
			board[bottom][col] = 0
			bottom--
		}
	}

	if bubblesBurst {
		return candyCrush(board)
	} else {
		return board
	}
}
