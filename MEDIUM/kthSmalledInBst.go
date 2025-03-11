package medium

var count int

func kthSmallest(root *TreeNode, k int) int {
	count = k
	return kthSmallestHelper(root).Val
}

func kthSmallestHelper(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := kthSmallestHelper(root.Left)
	count--
	if count == 0 {
		return root
	}
	right := kthSmallestHelper(root.Right)

	if left != nil {
		return left
	}
	return right

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
