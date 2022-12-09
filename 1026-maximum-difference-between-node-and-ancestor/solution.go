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

func traverse(root *TreeNode, max int, min int, m int) int {
	d := diff(max, min)
	if root == nil {
		return d
	}

	if root.Val < min {
		min = root.Val
	}

	if root.Val > max {
		max = root.Val
	}

	d = diff(max, min)

	left := traverse(root.Left, max, min, maximum(d, m))
	right := traverse(root.Right, max, min, maximum(d, m))
	max = maximum(maximum(left, right), m)

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
