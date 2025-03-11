package easy

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	leftTree := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(leftTree)
	return root
}
