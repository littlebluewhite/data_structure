package main

import "fmt"

var vertices = []string{
	"a", "b", "c", "d", "e",
}

type Node struct {
	value     string
	edgesList []*Node
}

func (n *Node) Connect(n2 *Node) {
	n.edgesList = append(n.edgesList, n2)
	n2.edgesList = append(n2.edgesList, n)
}

func (n *Node) GetAdjacentNodes() (adjacentNodes []string) {
	for _, node := range n.edgesList {
		adjacentNodes = append(adjacentNodes, node.value)
	}
	return
}

func (n *Node) IsConnect(n2 *Node) bool {
	for _, node := range n.edgesList {
		if node.value == n2.value {
			return true
		}
	}
	return false
}

type Graph struct {
	nodes []*Node
}

func NewGraph(nodes ...*Node) (g *Graph) {
	g = &Graph{nodes: nodes}
	return
}

func (g *Graph) Add(n *Node) {
	g.nodes = append(g.nodes, n)
}

func main() {
	a := &Node{value: "a"}
	b := &Node{value: "b"}
	c := &Node{value: "c"}
	d := &Node{value: "d"}
	e := &Node{value: "e"}

	graph := NewGraph(a, b, c, d, e)
	a.Connect(b)
	a.Connect(d)
	b.Connect(c)
	c.Connect(d)
	c.Connect(e)
	c.Connect(e)
	fmt.Printf("%v, %v\n", a, graph.nodes[0])
	fmt.Println(a == graph.nodes[0])

	fmt.Println(*graph.nodes[0])
	fmt.Println(a.GetAdjacentNodes())
	fmt.Println(a.IsConnect(c))
}

//time complexity to find adjacent nodes -> O(1)
//time complexity to check if two nodes are connected -> O(logv) if each adjacent row is sorted
//space complexity -> O(e)
