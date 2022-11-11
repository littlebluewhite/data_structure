package main

import "fmt"

type ASD struct {
	A int
	B int
}

func main() {
	a := []ASD{{1, 2}}
	fmt.Println(a)
	a[0].B = 3
	fmt.Println(a)
}
