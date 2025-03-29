package easy

import (
	"math"
)

var maxDia int = math.MinInt

func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}
	lf := dfs(root.Left)
	rt := dfs(root.Right)

	maxDia = max(maxDia, lf+rt)
	return maxDia

}

func dfs(root *TreeNode) int {

	if root == nil {
		return 0
	}

	lft := dfs(root.Left)
	rt := dfs(root.Right)

	maxDia = max(maxDia, lft+rt)

	return max(rt+1, lft+1)

}
