package main

import "fmt"

//input: numbers = [5, 5, 7, 7, 7, 11, 11, 11, 15, 15, 15, 15, 17, 17, 17, 17, 19], target = 22
//output: [[0, 12], [2, 8], [5, 5]]
//
//input: numbers = [1, 3, 3, 9, 9, 11, 11, 11, 11, 29, 30, 30, 31, 31, 39], target = 40
//output: [[0, 14], [3, 12], [5, 9]]

func solution3(numbers []int, target int) (result [][2]int) {
	f := 0
	l := len(numbers) - 1
	for f <= l {
		if numbers[l] == numbers[l-1] {
			l--
			continue
		}
		twoSum := numbers[f] + numbers[l]
		if twoSum < target {
			f++
		} else if twoSum > target {
			l--
		} else if twoSum == target {
			result = append(result, [2]int{f, l})
			f++
			l--
		}
	}
	return result
}

func main() {
	fmt.Println(solution3([]int{5, 5, 7, 7, 7, 11, 11, 11, 15, 15, 15, 15, 17, 17, 17, 17, 19}, 22))
	fmt.Println(solution3([]int{1, 3, 3, 9, 9, 11, 11, 11, 11, 29, 30, 30, 31, 31, 39}, 40))
}
