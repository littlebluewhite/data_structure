package common

import (
	"data_structure/util"
	"fmt"
)

type Node struct {
	Value     string
	EdgesList []GraphNode
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

type Graph struct {
	Nodes []GraphNode
}

type GraphNode interface {
	GetAdjacentNodes() []GraphNode
	GetValue() string
	Connect(node GraphNode)
	IsConnected(node GraphNode) bool
}

func (g *Graph) AddToGraph(n GraphNode) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) BreadthFirstTraversal(start, end GraphNode) {
	queue := []GraphNode{start}
	visited := map[string]struct{}{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		fmt.Println(c)
		if end.GetValue() == c.GetValue() {
			fmt.Println("found it")
			return
		}
		visited[c.GetValue()] = struct{}{}
		for _, node := range c.GetAdjacentNodes() {
			if _, ok := visited[node.GetValue()]; !ok {
				queue = append(queue, node.(GraphNode))
			}
		}
	}
}

func (g *Graph) GetShortestPath(start, end GraphNode) (path []string) {
	queue := []GraphNode{start}
	visited := map[string]GraphNode{start.GetValue(): nil}
	for len(queue) > 0 {
		c := queue[0]
		if c.GetValue() == end.GetValue() {
			fmt.Println("found it")
			path = g.reconstructPath(visited, end)
			return
		}
		queue = queue[1:]
		for _, node := range c.GetAdjacentNodes() {
			if _, ok := visited[node.GetValue()]; !ok {
				visited[node.GetValue()] = c
				queue = append(queue, node.(GraphNode))
			}
		}
	}
	return
}

func (g *Graph) reconstructPath(visited map[string]GraphNode, end GraphNode) (path []string) {
	path = append(path, end.GetValue())
	current := visited[end.GetValue()]
	for current != nil {
		path = append(path, current.GetValue())
		current = visited[current.GetValue()]
	}
	path = util.Reverse(path)
	return
}

func (g *Graph) DepthFirstTraversal(start, end GraphNode, visited map[string]struct{}) {
	fmt.Println(start)
	if start.GetValue() == end.GetValue() {
		fmt.Println("found it")
		return
	}
	visited[start.GetValue()] = struct{}{}
	for _, node := range start.GetAdjacentNodes() {
		if _, ok := visited[node.GetValue()]; !ok {
			g.DepthFirstTraversal(node.(GraphNode), end, visited)
		}
	}
}

func (g *Graph) TopologicalSort() {
	var topOrdering []string
	visited := map[string]struct{}{}
	for _, node := range g.Nodes {
		fmt.Println(visited)
		g.depthTraversal(node, visited, &topOrdering)
	}
	fmt.Println(util.Reverse(topOrdering))
}

func (g *Graph) depthTraversal(start GraphNode, visited map[string]struct{}, topOrdering *[]string) {
	if _, ok := visited[start.GetValue()]; ok {
		return
	}
	visited[start.GetValue()] = struct{}{}
	for _, node := range start.GetAdjacentNodes() {
		g.depthTraversal(node, visited, topOrdering)
	}

	*topOrdering = append(*topOrdering, start.GetValue())
}
