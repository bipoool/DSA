package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {

	queue := []*TreeNode{}
	start := -1

	queue = append(queue, root)

	res := []int{}

	for start < len(queue) {

		lenQ := len(queue)
		for i := start + 1; i < lenQ; i++ {
			currNode := queue[i]
			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
			if i == lenQ-1 {
				res = append(res, currNode.Val)
				start += lenQ
			}
		}

	}
	return res
}
