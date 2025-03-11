package medium

func levelOrder(root *TreeNode) [][]int {

	queue := make([]*TreeNode, 0)
	i := -1
	res := make([][]int, 0)

	queue = append(queue, root)
	i++

	for root != nil {

		currRow := make([]int, 0)
		lenQ := len(queue)
		for i < lenQ {
			node := queue[i]
			i++
			currRow = append(currRow, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		res = append(res, currRow)
	}
	return res
}
