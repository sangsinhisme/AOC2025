package Day7

import (
	"AOC2025/Utils"
	"fmt"
)

/*
--- Day 7: Laboratories ---
*/
const day = 7

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
	m := len(lines[0])
	dp := make([]bool, m)
	for i := range m {
		if lines[0][i] == 'S' {
			dp[i] = true
		}
	}
	ans := 0
	for _, line := range lines[1:] {
		for i, elem := range line {
			if elem == '^' && dp[i] {
				if i > 0 {
					dp[i-1] = true
				}
				if i < m-1 {
					dp[i+1] = true
				}
				dp[i] = false
				ans++
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
	m := len(lines[0])
	dp := make([]int, m)
	for i := range m {
		if lines[0][i] == 'S' {
			dp[i] = 1
		}
	}
	for _, line := range lines[1:] {
		for i, elem := range line {
			if elem == '^' && dp[i] != 0 {
				if i > 0 {
					dp[i-1] += dp[i]
				}
				if i < m-1 {
					dp[i+1] += dp[i]
				}
				dp[i] = 0
			}
		}
	}
	ans := 0
	for _, value := range dp {
		ans += value
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
