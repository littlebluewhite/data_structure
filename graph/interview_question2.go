//You are given a network of nodes as a graph. Some nodes are initially infected by MALWARE.
//The malware will spread from the initially infected nodes to all the nodes it is connected to, and it will continue
//to spread across the network accordingly.
//We have the opportunity to remove exactly ONE node from the initially infected list which will remove it and all its
//connections.
//Return the node that if removed, will minimize the number of nodes ultimately infected in this attack.

package main

import (
	"data_structure/common"
	"fmt"
)

func main() {
	network := [][]int{
		{0, 1, 0, 0},
		{1, 0, 1, 0},
		{0, 1, 0, 1},
		{0, 0, 1, 0},
	}
	infected := []int{0, 1}
	fmt.Println(minimizeMalwareSpread(network, infected))
}

func minimizeMalwareSpread(network [][]int, initiallyInfected []int) (n int) {
	worstMap := map[int]int{}
	bestSize := -len(network)
	infectedMap := map[int]struct{}{}
	data := make([]int, len(network))
	for i := 0; i < len(network); i++ {
		data[i] = -1
	}
	d := common.DisjointSet{Data: data}
	for _, n := range initiallyInfected {
		infectedMap[n] = struct{}{}
	}
	for n, i := range network {
		if _, ok := infectedMap[n]; ok {
			continue
		} else {
			for n2, j := range i {
				if _, ok := infectedMap[n2]; ok {
					continue
				} else {
					if j == 1 {
						d.Union(n, n2)
					}
				}
			}
		}
	}
	for _, deleteNode := range initiallyInfected {
		copyData := make([]int, len(d.GetData()))
		copy(copyData, d.GetData())
		newD := common.DisjointSet{Data: copyData}
		for _, infectedN := range initiallyInfected {
			if deleteNode == infectedN {
				continue
			} else {
				for _, infectedNode := range initiallyInfected {
					if infectedNode == deleteNode {
						continue
					} else {
						for i, value := range network[infectedNode] {
							if i == deleteNode {
								continue
							} else {
								if value == 1 {
									newD.Union(infectedNode, i)
								}
							}
						}
					}
				}
			}
		}
		fmt.Println(newD.GetData())
		worstSize := 0
		for _, infectedN := range initiallyInfected {
			if newD.GetData()[infectedN] < worstSize {
				worstSize = newD.GetData()[infectedN]
			}
		}
		worstMap[worstSize] = deleteNode
		if worstSize > bestSize {
			bestSize = worstSize
		}
	}
	return worstMap[bestSize]
}
