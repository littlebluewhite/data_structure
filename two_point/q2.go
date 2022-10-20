package main

import "fmt"

//input: numbers = [1, 2, 5, 7, 11, 15, 17, 29, 31], target = 22
//output: [[2, 6], [3, 5], [4, 4]]
//
//input: numbers = [1, 3, 6, 9, 10, 11, 13, 22, 28, 29, 31, 39], target = 40
//output: [[0, 11], [3, 10], [5, 9]]

func solution2(numbers []int, target int) (result [][2]int) {
	f := 0
	l := len(numbers) - 1
	for f <= l {
		twoSum := numbers[f] + numbers[l]
		if twoSum < target {
			f += 1
		} else if twoSum > target {
			l -= 1
		} else if twoSum == target {
			result = append(result, [2]int{f, l})
			f += 1
			l -= 1
		}
	}
	return
}

func main() {
	fmt.Println(solution2([]int{1, 2, 5, 7, 11, 15, 17, 29, 31}, 22))
	fmt.Println(solution2([]int{1, 3, 6, 9, 10, 11, 13, 22, 28, 29, 31, 39}, 40))
}
