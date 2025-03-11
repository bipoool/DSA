package easy

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if helper(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Left, subRoot)

}

func helper(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val {
		return helper(root.Left, subRoot.Left) && helper(root.Left, subRoot.Left)
	}

	return false

}
