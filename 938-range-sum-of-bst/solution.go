package rangesumofbst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	return sumTree(root, low, high)
}

func sumTree(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var h, l, r int
	if root.Val >= low && root.Val <= high {
		h = root.Val
	}

	if root.Left != nil && root.Val >= low {
		l = sumTree(root.Left, low, high)
	}

	if root.Right != nil && root.Val <= high {
		r = sumTree(root.Right, low, high)
	}

	return h + l + r
}
