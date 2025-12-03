package Day2

import (
	"AOC2025/Utils"
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 2: Gift Shop ---
*/
const day = 2

func helper(input string) bool {
	n := len(input)
	for i := 1; i <= n/2; i++ {
		match := input[:i]
		if n%i != 0 {
			continue
		}
		target := n / i
		matchPattern := 0
		for j := i; j <= n-i; j = j + i {
			if match == input[j:j+i] {
				matchPattern++
			} else {
				break
			}
		}
		if matchPattern == target-1 {
			return true
		}
	}
	return false
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
	lines = strings.Split(lines[0], ",")
	ans := 0
	for _, input := range lines {
		split := strings.Split(input, "-")
		start, end := split[0], split[1]
		startCov, _ := strconv.Atoi(start)
		endCov, _ := strconv.Atoi(end)
		for i := startCov; i <= endCov; i++ {
			stringI := strconv.Itoa(i)
			if len(stringI)%2 == 0 {
				half := len(stringI) / 2
				if stringI[:half] == stringI[half:] {
					fmt.Println(stringI)
					ans += i
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
	lines = strings.Split(lines[0], ",")
	ans := 0
	for _, input := range lines {
		split := strings.Split(input, "-")
		start, end := split[0], split[1]
		startCov, _ := strconv.Atoi(start)
		endCov, _ := strconv.Atoi(end)
		for i := startCov; i <= endCov; i++ {
			stringI := strconv.Itoa(i)
			if helper(stringI) {
				ans += i
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
