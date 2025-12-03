package Day1

import (
	"AOC2025/Utils"
	"fmt"
	"strconv"
)

/*
--- Day 1: Secret Entrance ---
*/
const day = 1

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
	dial := 50
	ans := 0
	fmt.Println("The dial starts by pointing at ", dial)
	for _, line := range lines {
		numRote, _ := strconv.Atoi(line[1:])
		numRote = numRote % 100
		if line[0] == 'L' {
			dial -= numRote
			if dial < 0 {
				dial += 100
			}
		} else {
			dial = (dial + numRote) % 100
		}
		if dial == 0 {
			ans += 1
		}
		fmt.Println("The dial is rotated " + line + " to point at " + strconv.Itoa(dial))
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
	dial := 50
	ans := 0
	fmt.Println("The dial starts by pointing at ", dial)
	for _, line := range lines {
		rotation, _ := strconv.Atoi(line[1:])
		direction := line[0]
		turns := rotation / 100
		rotate := rotation % 100
		ans += turns
		if direction == 'R' {
			if dial+rotate >= 100 {
				ans++
			}
		} else {
			if dial > 0 && (dial-rotate) <= 0 {
				ans++
			}
		}
		if direction == 'L' {
			dial = (dial - rotate + 100) % 100
		} else {
			dial = (dial + rotate) % 100
		}
	}
	fmt.Println("The dial end at " + strconv.Itoa(dial))
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
