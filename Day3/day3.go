package Day3

import (
	"AOC2025/Utils"
	"fmt"
)

/*
--- Day 3: Lobby ---
*/
const day = 3

func Part1(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
	}
	ans := 0
	for _, line := range lines {
		arr := make([]int, len(line))
		for i, char := range line {
			arr[i] = int(char - '0')
		}
		highest, idx := arr[0], 0
		for i, elem := range arr {
			if elem > highest {
				highest = elem
				idx = i
			}
		}
		if idx != len(arr)-1 {
			secondLeft := arr[idx+1]
			for i := idx + 1; i < len(arr); i++ {
				if arr[i] > secondLeft {
					secondLeft = arr[i]
				}
			}
			fmt.Printf("Joltagess for %s is %d", line, highest*10+secondLeft)
			ans += highest*10 + secondLeft
		} else {
			secondLeft := arr[0]
			for i := 0; i < idx; i++ {
				if arr[i] > secondLeft {
					secondLeft = arr[i]
				}
			}
			fmt.Printf("Joltagess for %s is %d", line, secondLeft*10+highest)
			ans += secondLeft*10 + highest
		}
		fmt.Println()
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, ans)
	}
}

func helper(arr []int, n, left, pivot int) (int, int) {
	value, nextPivot := arr[pivot], pivot
	for i := pivot; i <= n-left; i++ {
		if arr[i] > value {
			value = arr[i]
			nextPivot = i
		}
	}
	return value, nextPivot
}

func Part2(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
	}
	ans := 0
	for _, line := range lines {
		n := len(line)
		arr := make([]int, n)
		for i, char := range line {
			arr[i] = int(char - '0')
		}
		maxFind := 0
		possible, pivot := 0, 0
		for i := 0; i < 12; i++ {
			possible, pivot = helper(arr, n, 12-i, pivot)
			pivot++
			maxFind = maxFind*10 + possible
		}
		fmt.Printf("Value for %s is %d", line, maxFind)
		ans += maxFind
		fmt.Println()
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
