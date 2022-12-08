package leafsimilartrees

import (
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	res1 := []int{}
	res2 := []int{}
	traverse(root1, &res1)
	traverse(root2, &res2)

	return reflect.DeepEqual(res1, res2)
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	isLeaf := root.Left == nil && root.Right == nil
	if isLeaf {
		*res = append(*res, root.Val)
		return
	}

	traverse(root.Left, res)
	traverse(root.Right, res)
}
