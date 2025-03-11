package medium

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	rootIndex := findIndex(inorder, preorder[0])
	if rootIndex > 0 {
		root.Left = buildTree(preorder[1:rootIndex+1], inorder[:rootIndex])
	}
	if rootIndex < len(preorder)-1 {
		root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	}

	return root
}

func findIndex(arr []int, target int) int {
	for i := range arr {
		if target == arr[i] {
			return i
		}
	}
	return -1
}
