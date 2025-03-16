package medium

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	hMap := map[int]*Node{}
	return helper2(node, hMap)
}

func helper2(node *Node, hMap map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := hMap[node.Val]; ok {
		return v
	}
	res := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, 0),
	}
	hMap[node.Val] = res
	for _, n := range node.Neighbors {
		res.Neighbors = append(res.Neighbors, helper2(n, hMap))
	}
	return res
}
