package maximumdifferencebetweennodeandancestor

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return traverse(root, root.Val, root.Val, 0)
}

func traverse(root *TreeNode, max int, min int, currMax int) int {
	if root == nil {
		return currMax
	}

	if root.Val < min {
		min = root.Val
	}

	if root.Val > max {
		max = root.Val
	}

	d := diff(max, min)
	currMax = maximum(d, currMax)

	left := traverse(root.Left, max, min, currMax)
	right := traverse(root.Right, max, min, currMax)
	max = maximum(maximum(left, right), currMax)

	return max
}

func diff(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

func maximum(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
