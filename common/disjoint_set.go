package common

import "fmt"

type DisjointSet struct {
	Data []int
}

// [-1, -1, -1, -1]

func (d *DisjointSet) GetData() []int {
	return d.Data
}

func (d *DisjointSet) Find(node int) int {
	if d.Data[node] < 0 {
		return node
	} else {
		return d.Find(d.Data[node])
	}
}

func (d *DisjointSet) Union(n1, n2 int) {
	p1 := d.Find(n1)
	p2 := d.Find(n2)
	if p1 == p2 {
		fmt.Println("same parents!")
	} else {
		d.Data[p1] += d.Data[p2]
		d.Data[p2] = p1
	}
}
