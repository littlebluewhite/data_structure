package main

import "fmt"

func main() {
	fmt.Printf("type: %[1]T, %[1]q, %[1]c, %[1]v\n", []rune{'a', 'b', 'c'})
	fmt.Println(5 / 3)
}
