package common

import (
	"data_structure/util"
	"fmt"
)

type Node struct {
	Value     string
	EdgesList []*Node
}

func (n *Node) Connect(n2 *Node) {
	n.EdgesList = append(n.EdgesList, n2)
	n2.EdgesList = append(n2.EdgesList, n)
}

func (n *Node) GetAdjacentNodes() []*Node {
	return n.EdgesList
}

func (n *Node) IsConnected(n2 *Node) bool {
	for _, node := range n.EdgesList {
		if node == n2 {
			return true
		}
	}
	return false
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) AddToGraph(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) BreadthFirstTraversal(start, end *Node) {
	queue := []*Node{start}
	visited := map[string]struct{}{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		fmt.Println(c)
		if end.Value == c.Value {
			fmt.Println("found it")
			return
		}
		visited[c.Value] = struct{}{}
		for _, node := range c.EdgesList {
			if _, ok := visited[node.Value]; !ok {
				queue = append(queue, node)
			}
		}
	}
}

func (g *Graph) GetShortestPath(start, end *Node) (path []string) {
	queue := []*Node{start}
	visited := map[string]*Node{start.Value: nil}
	for len(queue) > 0 {
		c := queue[0]
		if c.Value == end.Value {
			fmt.Println("found it")
			path = g.reconstructPath(visited, end)
			return
		}
		queue = queue[1:]
		for _, node := range c.EdgesList {
			if _, ok := visited[node.Value]; !ok {
				visited[node.Value] = c
				queue = append(queue, node)
			}
		}
	}
	return
}

func (g *Graph) reconstructPath(visited map[string]*Node, end *Node) (path []string) {
	path = append(path, end.Value)
	current := visited[end.Value]
	for current != nil {
		path = append(path, current.Value)
		current = visited[current.Value]
	}
	path = util.Reverse(path)
	return
}
