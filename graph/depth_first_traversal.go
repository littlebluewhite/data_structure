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
	HKG := &common.Node{Value: "HKG"}
	g := &common.Graph{Nodes: []common.GraphNode{DFW, JFK, LAX, HNL, SAN, EWR, BOS, MIA, MCO, PBI, HKG}}
	DFW.Connect(LAX)
	DFW.Connect(JFK)
	LAX.Connect(HNL)
	LAX.Connect(EWR)
	LAX.Connect(SAN)
	SAN.Connect(HKG)
	JFK.Connect(BOS)
	JFK.Connect(MIA)
	MIA.Connect(MCO)
	MIA.Connect(PBI)
	MCO.Connect(PBI)
	g.DepthFirstTraversal(DFW, HKG, map[string]struct{}{})
}

//Traverses deeply into the data structure by exploring all nodes on a branch going forword until you reach
//the node you are searching for or a dead end.

//Backtrack up the branch until you reach the new branch to explore

//Uses a stack data structure

//time complexity O(v+e) where v is the number of vertices and e is the number of edges
//space complexity is O(v) because the visited array will at most store each vertex if we traverse
//the entire graph
