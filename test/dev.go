package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	fmt.Printf("type: %[1]T, %[1]q, %[1]c, %[1]v\n", []rune{'a', 'b', 'c'})
	a := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	b := a[3:6]
	c := a[6:9]
	fmt.Printf("%v, %v, %v, %v\n", len(b), cap(b), len(c), cap(c))
	fmt.Println(c)
	d := [...]int{0, 1, 2, 3, 4, 5}
	reverse(d[:])
	fmt.Println(d)
	e := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v, %v\n", len(e), cap(e))
	reverse(e[:2])
	fmt.Println(e)
	reverse(e[2:])
	fmt.Println(e)
	reverse(e)
	fmt.Println(e)
}
