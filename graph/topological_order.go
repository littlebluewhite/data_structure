package main

import "data_structure/common"

func main() {
	a := &common.DirectNode{Value: "a"}
	b := &common.DirectNode{Value: "b"}
	c := &common.DirectNode{Value: "c"}
	d := &common.DirectNode{Value: "d"}
	e := &common.DirectNode{Value: "e"}
	f := &common.DirectNode{Value: "f"}

	g := &common.Graph{Nodes: []common.GraphNode{a, b, c, d, e, f}}

	a.Connect(c)
	a.Connect(b)
	b.Connect(d)
	d.Connect(f)
	e.Connect(f)
	e.Connect(c)

	g.TopologicalSort()
}
