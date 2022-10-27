package common

type Node struct {
	Value     string
	EdgesList []GraphNode
	Color     string
}

func (n *Node) Connect(n2 GraphNode) {
	n3 := n2.(*Node)
	n.EdgesList = append(n.EdgesList, n2)
	n3.EdgesList = append(n3.EdgesList, n)
}

func (n *Node) GetValue() string {
	return n.Value
}

func (n *Node) GetAdjacentNodes() []GraphNode {
	return n.EdgesList
}

func (n *Node) IsConnected(n2 GraphNode) bool {
	for _, node := range n.EdgesList {
		if node == n2 {
			return true
		}
	}
	return false
}
