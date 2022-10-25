package main

import (
	"data_structure/common"
	"fmt"
)

func main() {
	Hannah := &common.Node{Value: "Hannah"}
	Mary := &common.Node{Value: "Mary"}
	Mel := &common.Node{Value: "Mel"}
	Max := &common.Node{Value: "Max"}
	Nis := &common.Node{Value: "Nis"}
	Dan := &common.Node{Value: "Dan"}
	Chris := &common.Node{Value: "Chris"}
	Yair := &common.Node{Value: "Yair"}
	Nicolette := &common.Node{Value: "Nicolette"}
	Mabel := &common.Node{Value: "Mabel"}
	Liz := &common.Node{Value: "Liz"}
	g := &common.Graph{Nodes: []*common.Node{
		Hannah, Mary, Mel, Max, Nis, Dan, Chris, Yair, Nicolette, Mabel, Liz}}
	Hannah.Connect(Mary)
	Hannah.Connect(Max)
	Hannah.Connect(Mel)
	Hannah.Connect(Nis)
	Hannah.Connect(Liz)
	Max.Connect(Mel)
	Nis.Connect(Dan)
	Nis.Connect(Yair)
	Nis.Connect(Chris)
	Yair.Connect(Chris)
	Yair.Connect(Mabel)
	Yair.Connect(Liz)
	Mabel.Connect(Liz)
	Chris.Connect(Nicolette)
	fmt.Println(g.GetShortestPath(Mel, Yair))
}
