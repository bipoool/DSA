package hard

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}
	queue := make([]*TreeNode, 0)
	i := -1

	queue = append(queue, root)
	i++
	serialisedArr := make([]string, 0)
	for i < len(queue) {

		currArr := make([]string, 0)
		currLen := len(queue)

		for i < currLen {
			node := queue[i]
			i++

			if node == nil {
				currArr = append(currArr, "nil")
				continue
			}

			currArr = append(currArr, fmt.Sprintf("%d", node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)

		}
		currArrStr := strings.Join(currArr, ",")
		serialisedArr = append(serialisedArr, currArrStr)

	}
	fmt.Println(strings.Join(serialisedArr, "->"))
	return strings.Join(serialisedArr, "->")

}

// Deserializes your encoded data to tree.
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {

	var bfsArr [][]string
	bfsTempArr := strings.Split(data, "->")

	for _, arr := range bfsTempArr {
		bfsArr = append(bfsArr, strings.Split(arr, ","))
	}

	queue := make([]*TreeNode, 0)
	i := -1

	root := &TreeNode{}
	queue = append(queue, root)
	i++

	for _, arr := range bfsArr {
		for _, val := range arr {
			node := queue[i]
			i++
			if val != "nil" {
				intVal, _ := strconv.Atoi(val)
				node.Val = int(intVal)
			} else {
				node.Val = math.MinInt
				continue
			}
			node.Left = &TreeNode{}
			node.Right = &TreeNode{}
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return deleteNullNodes(root)
}

func deleteNullNodes(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	if root.Left != nil && root.Left.Val == math.MinInt {
		root.Left = nil
	}

	if root.Right != nil && root.Right.Val == math.MinInt {
		root.Right = nil
	}

	if root.Left != nil {
		deleteNullNodes(root.Left)
	}
	if root.Right != nil {
		deleteNullNodes(root.Right)
	}

	return root

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
