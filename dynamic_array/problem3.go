package main

import (
	"data_structure/common"
	"fmt"
)

//message = [
//	'e', 'l', 's', 'e', ' ',
//	'o', 'r', ' ',
//	'm', 'o', 'n', 'e', 'y', ' ',
//	'a', 'l', 'l', ' ',
//	'u', 's', ' ',
//	'g', 'i', 'v', 'e'
//]

// message2 = [
//	'o', 'u', 't', 's', 'i', 'd', 'e', ' ',
//	'p', 'l', 'a', 'y', 'i', 'n', 'g', ' ',
//	'l', 'o', 'v', 'e', 's', ' ',
//	'd', 'o', 'g', ' ',
//	't', 'h', 'e'
// ]

func main() {
	m1 := []rune{
		'e', 'l', 's', 'e', ' ',
		'o', 'r', ' ',
		'm', 'o', 'n', 'e', 'y', ' ',
		'a', 'l', 'l', ' ',
		'u', 's', ' ',
		'g', 'i', 'v', 'e',
	}
	m2 := []rune{
		'o', 'u', 't', 's', 'i', 'd', 'e', ' ',
		'p', 'l', 'a', 'y', 'i', 'n', 'g', ' ',
		'l', 'o', 'v', 'e', 's', ' ',
		'd', 'o', 'g', ' ',
		't', 'h', 'e',
	}
	fmt.Println(string(common.ReverseWords(m1)))
	fmt.Println(string(common.ReverseWords(m2)))
}
