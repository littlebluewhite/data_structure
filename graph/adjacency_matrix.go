package main

import "fmt"

var vertices = []string{"A", "B", "C", "D", "E"}

var edges = [][]string{
	{"A", "B"},
	{"A", "D"},
	{"B", "C"},
	{"C", "D"},
	{"C", "E"},
	{"D", "E"},
}

func main() {
	fmt.Println(vertices)
	fmt.Println(edges)
	fmt.Println(findAdjacentNodes("C"))
	fmt.Println(isConnected("A", "E"))
}

func findAdjacentNodes(node string) (adjacentNodes []string) {
	for _, edge := range edges {
		for idx, s := range edge {
			if s == node {
				switch idx {
				case 0:
					adjacentNodes = append(adjacentNodes, edge[1])
				case 1:
					adjacentNodes = append(adjacentNodes, edge[0])
				default:
				}
			}
		}
	}
	return
}

func isConnected(node1, node2 string) bool {
	for _, edge := range edges {
		if edge[0] == node1 && edge[1] == node2 {
			return true
		} else if edge[0] == node2 && edge[1] == node1 {
			return true
		}
	}
	return false
}
