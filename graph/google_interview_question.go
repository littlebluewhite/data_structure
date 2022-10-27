//You are tasked with splitting the 7th grade class into two teams for a soccer game.
//Given a list of students and their "enemies" aka the kids they can't be on th same team as,
//determine whether it is possible to split the class in two in such a way that no child is on the same
//team as any of their enemies. if possible, return the two teams.

//David: [Lucy, Jose, Chris]
//Lucy: [David, Brian, Emily]
//Emily: [Lucy, Jack]
//Jose: [David, Paul]
//Paul: [Jose, Chris]
//Chris: [Paul, David, Brian]
//Brian: [Lucy, Chris, Jack]
//Jack: [Brian, Emily]

package main

import "data_structure/common"

func main() {
	David := &common.Node{Value: "David"}
	Lucy := &common.Node{Value: "Lucy"}
	Emily := &common.Node{Value: "Emily"}
	Jose := &common.Node{Value: "Jose"}
	Paul := &common.Node{Value: "Paul"}
	Chris := &common.Node{Value: "Chris"}
	Brian := &common.Node{Value: "Brian"}
	Jack := &common.Node{Value: "Jack"}

	g := common.Graph{Nodes: []common.GraphNode{David, Lucy, Emily, Jose, Paul, Chris, Brian, Jack}}

	David.Connect(Lucy)
	David.Connect(Jose)
	David.Connect(Chris)
	Lucy.Connect(Brian)
	Lucy.Connect(Emily)
	Emily.Connect(Jack)
	Jose.Connect(Paul)
	Paul.Connect(Chris)
	Chris.Connect(Brian)
	Brian.Connect(Jack)
	//Jack.Connect(Chris)

	g.IsBipartite()
}
