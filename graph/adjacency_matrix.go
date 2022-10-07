package main

import "fmt"

var vertices2 = []string{"A", "B", "C", "D", "E"}

var vertexIdx = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
}

var adjacencyMatrix = [][]int{
	{0, 1, 0, 1, 0},
	{1, 0, 1, 0, 0},
	{0, 1, 0, 1, 1},
	{1, 0, 1, 0, 1},
	{0, 0, 1, 1, 0},
}

func main() {
	fmt.Println(findAdjacentNodes2("C"))
	fmt.Println(isConnect2("E", "C"))
}

func findAdjacentNodes2(node string) (adjacencyNodes []string) {
	idx := vertexIdx[node]
	for i, b := range adjacencyMatrix[idx] {
		if b == 1 {
			adjacencyNodes = append(adjacencyNodes, vertices2[i])
		}
	}
	return
}

func isConnect2(node1, node2 string) bool {
	idx1 := vertexIdx[node1]
	idx2 := vertexIdx[node2]
	if adjacencyMatrix[idx1][idx2] == 1 {
		return true
	} else {
		return false
	}
}

// time complexity O(v) to find adjacent nodes
// time complexity O(1) to check if two nodes are connected
// space complexity O(v^2)
