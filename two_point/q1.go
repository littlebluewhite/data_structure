package main

import "fmt"

//input: numbers = [1, 2, 4, 7, 11, 15, 17, 29, 31], target = 22
//output: [3, 5]
//
//input: numbers = [9, 11, 12, 22, 28, 29], target = 33
//output: [1, 3]

func solution(numbers []int, target int) [2]int {
	f := 0
	l := len(numbers) - 1
	for f <= l {
		twoSum := numbers[f] + numbers[l]
		if twoSum < target {
			f += 1
		} else if twoSum > target {
			l -= 1
		} else if twoSum == target {
			return [2]int{f, l}
		}
	}
	return [2]int{-1, -1}
}

func main() {
	fmt.Println(solution([]int{1, 2, 4, 7, 11, 15, 17, 29, 31}, 22))
	fmt.Println(solution([]int{9, 11, 12, 22, 28, 29}, 33))
}
