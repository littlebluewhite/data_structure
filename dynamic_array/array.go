package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Printf("%v, %T, %p\n", a, a, &a)
	fmt.Printf("%v, %T, %p\n", b, b, &b)
	var c *[]int = &b
	fmt.Printf("%v, %T, %p\n", c, c, c)
	b = append(b, 4)
	a = append(b, 6)
	fmt.Printf("%v, %T, %p\n", a, a, &a)
	fmt.Printf("%v, %T, %p\n", b, b, &b)
	fmt.Printf("%v, %T, %p\n", c, c, c)

}
