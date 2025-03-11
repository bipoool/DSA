package ds

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrayToTree(arr []int) *TreeNode {
	arrI := 0
	i := 0
	root := &TreeNode{
		Val: arr[0],
	}
	curr := root

	queue := make([]*TreeNode, len(arr))
	queue[i] = root
	i++

	for arrI < len(arr)-2 {
		curr = queue[i-1]
		i--
		leftVal := arr[arrI+1]
		rightVal := arr[arrI+2]

		if leftVal != math.MaxInt {
			curr.Left = &TreeNode{
				Val: leftVal,
			}
			queue[i] = curr.Left
			i++
		}
		if rightVal != math.MaxInt {
			curr.Right = &TreeNode{
				Val: rightVal,
			}
			queue[i] = curr.Right
			i++
		}
		arrI += 2
	}
	return root
}
