package Day6

import (
	"AOC2025/Utils"
	"fmt"
	"regexp"
	"strconv"
)

/*
--- Day 6: Trash Compactor ---
*/
const day = 6

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
	mathRegex := regexp.MustCompile(`[*+]`)
	matches := mathRegex.FindAllString(lines[n-1], -1)

	numElem := len(matches)
	mathArr := make([]bool, numElem)
	for i, match := range matches {
		mathArr[i] = match == "*"
	}
	digitalRegex := regexp.MustCompile(`\d+`)
	ansArr := make([]int64, numElem)
	for _, line := range lines[:n-1] {
		matchDigits := digitalRegex.FindAllString(line, -1)
		for i, match := range matchDigits {
			value, _ := strconv.Atoi(match)
			if ansArr[i] == 0 {
				ansArr[i] = int64(value)
			} else if mathArr[i] {
				ansArr[i] *= int64(value)
			} else {
				ansArr[i] += int64(value)
			}
		}
		fmt.Println(ansArr)
	}
	ans := int64(0)
	for _, value := range ansArr {
		ans += value
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.SubmitString(day, 1, strconv.FormatInt(ans, 10))
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
	inverse := make([]string, m)
	for _, line := range lines {
		for i, char := range line {
			inverse[i] = inverse[i] + string(char)
		}
	}
	ans := int64(0)
	previous := uint8('+')
	eachMath := int64(0)
	for _, line := range inverse {
		if line[n-1] != ' ' {
			previous = line[n-1]
		}
		digitalRegex := regexp.MustCompile(`\d+`)
		if !digitalRegex.MatchString(line) {
			ans += eachMath
			eachMath = 0
		}
		matchDigits := digitalRegex.FindString(line)
		value, _ := strconv.Atoi(matchDigits)
		if eachMath == 0 {
			eachMath = int64(value)
		} else if previous == '+' {
			eachMath += int64(value)
		} else {
			eachMath *= int64(value)
		}
	}
	ans += eachMath
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.SubmitString(day, 2, strconv.FormatInt(ans, 10))
	}
}
