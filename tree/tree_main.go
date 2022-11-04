package main

import (
	"data_structure/common"
	"fmt"
)

func main() {
	a := common.NewTreeNode("a")
	b := common.NewTreeNode("b")
	c := common.NewTreeNode("c")
	d := common.NewTreeNode("d")
	e := common.NewTreeNode("e")
	f := common.NewTreeNode("f")
	a.SetLeft(b)
	a.SetRight(c)
	b.SetLeft(d)
	b.SetRight(e)
	c.SetRight(f)

	a.BreadthFirstPrint()
	fmt.Println("-----------------------------")
	fmt.Println(a.BreadthFirstSearch("c"))
	fmt.Println(a.BreadthFirstSearch("e"))
	fmt.Println(a.BreadthFirstSearch("h"))
	a.DepthFirstPrint()
	fmt.Println("-----------------------------")
	common.DepthFirstPrint2[string](a)
	fmt.Println("-----------------------------")
	a.PreOrder()
	fmt.Println("-----------------------------")
	a.InOrder()
	fmt.Println("-----------------------------")
	a.PostOrder()
	fmt.Println("-----------------------------")
}
