package leetcode

func main() {

}

func numberOfSubmatrices(grid [][]byte) int {
	n, m := len(grid), len(grid[0])

	xCount, yCount := make([]int, m), make([]int, m)
	ans := 0
	for i := range n {
		x, y := 0, 0
		for j := range m {
			switch grid[i][j] {
			case 'X':
				x++
			case 'Y':
				y++
			}
			xCount[j] += x
			yCount[j] += y
			if xCount[j] == yCount[j] && xCount[j] > 0 {
				ans++
			}
		}
	}
	return ans
}
