package medium

// Save sum at every step in pre-order traversal
// if compliment (currSum - target) is available add it

var h map[int]int

func pathSum(root *TreeNode, t int) int {
	h = map[int]int{}
	h[0] = 1
	return dfspathSum(root, 0, t)
}

func dfspathSum(root *TreeNode, cs int, t int) int {
	if root == nil {
		return 0
	}
	cs += root.Val

	count := 0
	if h[cs-t] != 0 {
		count += h[cs-t]
	}
	h[cs]++
	count += dfspathSum(root.Left, cs, t)
	count += dfspathSum(root.Right, cs, t)
	h[cs]--

	return count
}
