package Day4

import (
	"AOC2025/Utils"
	"fmt"
)

/*
--- Day 4: Printing Department ---
*/
const day = 4

func helper(grid [][]int32, n, m, i, j int) bool {
	count := 0
	for ver := -1; ver <= 1; ver++ {
		for hor := -1; hor <= 1; hor++ {
			nextI, nextJ := i+ver, j+hor
			if nextI >= 0 && nextI < n && nextJ >= 0 && nextJ < m {
				if grid[nextI][nextJ] == '@' {
					count++
				}
			}
		}
	}
	return count-1 < 4
}
func Part1(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/Day%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/Day%v.txt", day))
	}
	n := len(lines)
	m := len(lines[0])
	grid := make([][]int32, n)
	for i := range n {
		grid[i] = make([]int32, m)
	}
	for i, line := range lines {
		for j, char := range line {
			grid[i][j] = char
		}
	}
	ans := 0
	for i := range n {
		for j := range m {
			if grid[i][j] != '.' {
				if helper(grid, n, m, i, j) {
					ans += 1
				}
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, ans)
	}
}

func Part2(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/Day%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/Day%v.txt", day))
	}
	n := len(lines)
	m := len(lines[0])
	grid := make([][]int32, n)
	for i := range n {
		grid[i] = make([]int32, m)
	}
	for i, line := range lines {
		for j, char := range line {
			grid[i][j] = char
		}
	}
	ans := 0
	round := 1
	var nextRemove [][]int
	for {
		for i := range n {
			for j := range m {
				if grid[i][j] != '.' {
					if helper(grid, n, m, i, j) {
						nextRemove = append(nextRemove, []int{i, j})
					}
				}
			}
		}
		if len(nextRemove) == 0 {
			break
		}
		for _, elem := range nextRemove {
			grid[elem[0]][elem[1]] = '.'
		}
		fmt.Printf("Round %d remove %d elem \n", round, len(nextRemove))
		ans += len(nextRemove)
		round++
		nextRemove = [][]int{}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
