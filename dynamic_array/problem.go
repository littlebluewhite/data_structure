package main

import "fmt"

// Given a string, write a function to check if it is a permutation of a palindrome.

func main() {
	fmt.Println(permutationOfPalindromeChecker1("level"))
	fmt.Println(permutationOfPalindromeChecker2("level"))
	fmt.Println(permutationOfPalindromeChecker1("on meoln no loenm"))
	fmt.Println(permutationOfPalindromeChecker2("on meoln no loenm"))
	fmt.Println(permutationOfPalindromeChecker1("waas tii tca aws"))
	fmt.Println(permutationOfPalindromeChecker2("waas tii tca aws"))
	fmt.Println(permutationOfPalindromeChecker1("hnnaha"))
	fmt.Println(permutationOfPalindromeChecker2("hnnaha"))
	fmt.Println(permutationOfPalindromeChecker1("elisabeth"))
	fmt.Println(permutationOfPalindromeChecker2("elisabeth"))
	fmt.Println(permutationOfPalindromeChecker1("becky"))
	fmt.Println(permutationOfPalindromeChecker2("becky"))
}

func isCharacter(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	} else {
		return false
	}
}

func permutationOfPalindromeChecker1(s string) bool {
	letterCount := map[rune]int{}
	for _, r := range s {
		if isCharacter(r) {
			if _, ok := letterCount[r]; ok {
				letterCount[r]++
			} else {
				letterCount[r] = 1
			}
		}
	}
	fmt.Println(letterCount)
	oddNumber := 0
	for _, val := range letterCount {
		if val%2 != 0 {
			oddNumber++
			if oddNumber > 1 {
				return false
			}
		}
	}
	return true
}

func permutationOfPalindromeChecker2(s string) bool {
	letterBool := map[rune]bool{}
	for _, r := range s {
		if isCharacter(r) {
			if _, ok := letterBool[r]; !ok {
				letterBool[r] = false
			} else {
				letterBool[r] = !letterBool[r]
			}
		}
	}
	fmt.Println(letterBool)
	oddNumber := 0
	for _, val := range letterBool {
		if !val {
			oddNumber++
			if oddNumber > 1 {
				return false
			}
		}
	}
	return true
}
