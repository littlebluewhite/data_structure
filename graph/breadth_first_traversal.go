package main

import "data_structure/common"

func main() {
	DFW := &common.Node{Value: "DFW"}
	JFK := &common.Node{Value: "JFK"}
	LAX := &common.Node{Value: "LAX"}
	HNL := &common.Node{Value: "HNL"}
	SAN := &common.Node{Value: "SAN"}
	EWR := &common.Node{Value: "EWR"}
	BOS := &common.Node{Value: "BOS"}
	MIA := &common.Node{Value: "MIA"}
	MCO := &common.Node{Value: "MCO"}
	PBI := &common.Node{Value: "PBI"}
	g := &common.Graph{Nodes: []*common.Node{DFW, JFK, LAX, HNL, SAN, EWR, BOS, MIA, MCO, PBI}}
	DFW.Connect(LAX)
	DFW.Connect(JFK)
	LAX.Connect(HNL)
	LAX.Connect(EWR)
	LAX.Connect(SAN)
	JFK.Connect(BOS)
	JFK.Connect(MIA)
	MIA.Connect(MCO)
	MIA.Connect(PBI)
	MCO.Connect(PBI)
	g.BreadthFirstTraversal(DFW, MIA)
}

//Traversal broad into the data structure by visiting neighbor common.Nodes before children common.Nodes
//Uses a queue data structure

//breadth first Search has a time complexity of O(v+e)where v is th numbers of vertices and e
//is the number of edges
//The space complexity is O(v)
