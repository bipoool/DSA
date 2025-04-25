package medium

// Save sum at every step in pre-order traversal
// if compliment (currSum - target) is available add it

var h map[int]int

func pathSum(root *TreeNode, t int) int {
	h = map[int]int{}
	h[0] = 1
	return dfs(root, 0, t)
}

func dfs(root *TreeNode, cs int, t int) int {
	if root == nil {
		return 0
	}
	cs += root.Val

	count := 0
	if h[cs-t] != 0 {
		count += h[cs-t]
	}
	h[cs]++
	count += dfs(root.Left, cs, t)
	count += dfs(root.Right, cs, t)
	h[cs]--

	return count
}
