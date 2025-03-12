package medium

import "math"

var maxSum int = math.MinInt

func maxPathSum(root *TreeNode) int {
	maxPathSumHelper(root)
	return maxSum
}

func maxPathSumHelper(root *TreeNode) int {

	if root == nil {
		return 0
	}

	leftSum := maxPathSumHelper(root)
	rightSum := maxPathSumHelper(root.Right)

	currMaxSum := math.Max(float64(leftSum)+float64(root.Val), float64(rightSum)+float64(root.Val))

	if maxSum < leftSum+rightSum+root.Val {
		maxSum = leftSum + rightSum + root.Val
	}

	if maxSum < int(currMaxSum) {
		maxSum = int(currMaxSum)
	}

	if currMaxSum < 0 {
		return 0
	}

	return int(currMaxSum)

}
