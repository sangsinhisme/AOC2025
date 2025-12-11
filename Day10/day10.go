package Day10

import (
	"AOC2025/Utils"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 10: Factory ---
*/
const day = 10

func similarity(target, state []int) bool {
	for i := 0; i < len(target); i++ {
		if target[i] != state[i]%2 {
			return false
		}
	}
	return true
}
func backtracking(target, state []int, index int, buttons [][]int, press int) int {
	if index == len(buttons) {
		if similarity(target, state) {
			return press
		}
		return math.MaxInt
	}
	if similarity(target, state) {
		return press
	}
	minWithout := backtracking(target, state, index+1, buttons, press)
	for _, button := range buttons[index] {
		state[button]++
	}
	minWith := backtracking(target, state, index+1, buttons, press+1)
	for _, button := range buttons[index] {
		state[button]--
	}
	return min(minWithout, minWith)
}

func solvePuLP(buttons [][]int, target []int) int {
	// Convert data to JSON
	data := map[string]interface{}{
		"buttons": buttons,
		"target":  target,
	}
	jsonData, _ := json.Marshal(data)

	// Create command with JSON as stdin
	cmd := exec.Command("/usr/bin/python3", "Day10/day10part2.py")
	cmd.Stdin = strings.NewReader(string(jsonData))

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error: %v\nOutput: %s", err, output)
	}

	result, _ := strconv.Atoi(strings.TrimSpace(string(output)))
	return result
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
	ans := 0
	for _, line := range lines {
		lightRegex := regexp.MustCompile(`\[(.+?)]`)
		lights := lightRegex.FindStringSubmatch(line)[1]

		buttonRegex := regexp.MustCompile(`\((.+?)\)`)
		buttonExtract := buttonRegex.FindAllStringSubmatch(line, -1)
		buttons := make([][]int, len(buttonExtract))
		for i, button := range buttonExtract {
			split := strings.Split(button[1], ",")
			for _, index := range split {
				indexInt, _ := strconv.Atoi(index)
				buttons[i] = append(buttons[i], indexInt)
			}
		}

		target := make([]int, len(lights))
		for i, light := range lights {
			if light == '#' {
				target[i] = 1
			}
		}
		ans += backtracking(target, make([]int, len(lights)), 0, buttons, 0)
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
	ans := 0
	for _, line := range lines {
		joltagesRegex := regexp.MustCompile(`\{(.+?)}`)
		joltages := joltagesRegex.FindStringSubmatch(line)[1]

		buttonRegex := regexp.MustCompile(`\((.+?)\)`)
		buttonExtract := buttonRegex.FindAllStringSubmatch(line, -1)
		buttons := make([][]int, len(buttonExtract))
		for i, button := range buttonExtract {
			split := strings.Split(button[1], ",")
			for _, index := range split {
				indexInt, _ := strconv.Atoi(index)
				buttons[i] = append(buttons[i], indexInt)
			}
		}

		split := strings.Split(joltages, ",")
		target := make([]int, len(split))
		for i, elem := range split {
			value, _ := strconv.Atoi(elem)
			target[i] = value
		}

		presses := solvePuLP(buttons, target)
		fmt.Println(target, buttons, presses)
		ans += presses
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
