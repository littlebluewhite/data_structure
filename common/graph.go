package common

import (
	"data_structure/util"
	"errors"
	"fmt"
)

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

func (g *Graph) IsBipartite() {
	for _, node := range g.Nodes {
		n := node.(*Node)
		if n.Color == "" {
			n.Color = "red"
			fmt.Println("isBipartite", n)
			e := g.assignLegalColoring(n)
			if e != nil {
				fmt.Println(e)
				return
			}
		}
	}
	fmt.Println("graph is bipartite")
	var red []string
	var blue []string
	for _, node := range g.Nodes {
		n := node.(*Node)
		if n.Color == "red" {
			red = append(red, n.GetValue())
		} else {
			blue = append(blue, n.GetValue())
		}
	}
	fmt.Println(red)
	fmt.Println(blue)
}

func (g *Graph) assignLegalColoring(start *Node) (err error) {
	fmt.Println("assign", start)
	childColor := g.getChildColor(start)
	for _, node := range start.GetAdjacentNodes() {
		n2 := node.(*Node)
		if n2.Color == "" {
			n2.Color = childColor
			e := g.assignLegalColoring(n2)
			if e != nil {
				err = e
				return e
			}
		} else {
			if n2.Color == childColor {
				continue
			} else {
				err = errors.New("graph is not bipartite")
				return
			}
		}
	}
	return err
}

func (g *Graph) getChildColor(parent *Node) string {
	if parent.Color == "red" {
		return "blue"
	} else if parent.Color == "blue" {
		return "red"
	} else {
		return ""
	}
}
