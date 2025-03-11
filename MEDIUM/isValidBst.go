package medium

func isValidBST(root *TreeNode) bool {
	arr := make([]int, 0)
	arr = dfs(root, arr)

	for i := range arr {
		if i+1 < len(arr) && arr[i] >= arr[i+1] {
			return false
		}
	}

	return true

}

func dfs(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = dfs(root.Left, arr)
	arr = append(arr, root.Val)
	arr = dfs(root.Right, arr)
	return arr
}
