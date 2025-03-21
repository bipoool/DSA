package medium

import (
	"math"
)

// @TBD Check this later
var commonNode *TreeNode = &TreeNode{
	Val: math.MinInt,
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	helperlowestCommonAncestor(root, p, q)
	return commonNode
}

func helperlowestCommonAncestor(root *TreeNode, node1 *TreeNode, node2 *TreeNode) bool {

	if root == nil {
		return false
	}

	var rootFlag bool = false

	if root == node1 || root == node2 {
		rootFlag = true
	}

	left := helperlowestCommonAncestor(root.Left, node1, node2)
	right := helperlowestCommonAncestor(root.Right, node1, node2)

	if rootFlag {
		if right || left {
			commonNode = root
		}
	} else if right && left {
		commonNode = root
	}

	return right || left || rootFlag
}
