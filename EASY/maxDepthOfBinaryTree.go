package easy

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	hl := maxDepth(root.Left)
	h2 := maxDepth(root.Right)

	if hl > h2 {
		return hl
	}
	return h2

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
