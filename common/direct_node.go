package common

type DirectNode struct {
	Value     string
	EdgesList []GraphNode
}

func (d *DirectNode) Connect(d2 GraphNode) {
	d.EdgesList = append(d.EdgesList, d2)
}

func (d *DirectNode) GetValue() string {
	return d.Value
}

func (d *DirectNode) GetAdjacentNodes() []GraphNode {
	return d.EdgesList
}

func (d *DirectNode) IsConnected(d2 GraphNode) bool {
	d2 = d2.(*DirectNode)
	for _, node := range d.EdgesList {
		if node == d2 {
			return true
		}
	}
	return false
}
