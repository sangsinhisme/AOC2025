package Day17

import (
	"AOC2025/Utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
--- Day 17: Chronospatial Computer ---
*/
const day = 17

func Part1(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("2024/Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("2024/Input/Day%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("2024/Input/Day%v.txt", day))
	}
	registerA, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	registerB, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	registerC, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])

	programStr := strings.Split(lines[4], ": ")[1]
	program := strings.Split(programStr, ",")
	var ans []string

	var getOperand func(value, opcode int) int
	getOperand = func(value, opcode int) int {
		if value >= 0 && value <= 3 || opcode == 1 || opcode == 3 {
			return value
		} else if value == 4 {
			return registerA
		} else if value == 5 {
			return registerB
		} else if value == 6 {
			return registerC
		} else {
			return 7
		}
	}

	var instruction func(opcode, value, pointer int) int
	instruction = func(opcode, value, pointer int) int {
		if opcode == 0 {
			registerA = registerA / int(math.Pow(2, float64(value)))
		} else if opcode == 1 {
			registerB = registerB ^ value
		} else if opcode == 2 {
			registerB = value % 8
		} else if opcode == 3 {
			if registerA != 0 {
				return value
			}
		} else if opcode == 4 {
			registerB = registerB ^ registerC
		} else if opcode == 5 {
			ans = append(ans, strconv.Itoa(value%8))
		} else if opcode == 6 {
			registerB = registerA / int(math.Pow(2, float64(value)))
		} else if opcode == 7 {
			registerC = registerA / int(math.Pow(2, float64(value)))
		}
		return pointer + 2
	}

	programInt := make([]int, len(program))
	for i, str := range program {
		programInt[i], _ = strconv.Atoi(str)
	}

	pointer := 0
	for pointer < len(programInt) {
		pointer = instruction(programInt[pointer], getOperand(programInt[pointer+1], programInt[pointer]), pointer)
	}
	resultString := strings.Join(ans, ",")
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.SubmitString(day, 1, resultString)
	}
}

func Part2(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("2024/Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("2024/Input/Day%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("2024/Input/Day%v.txt", day))
	}
	registerA, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	registerB, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	registerC, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])

	programStr := strings.Split(lines[4], ": ")[1]
	program := strings.Split(programStr, ",")
	var ans []string

	var getOperand func(value, opcode int) int
	getOperand = func(value, opcode int) int {
		if value >= 0 && value <= 3 || opcode == 1 || opcode == 3 {
			return value
		} else if value == 4 {
			return registerA
		} else if value == 5 {
			return registerB
		} else if value == 6 {
			return registerC
		} else {
			return 7
		}
	}

	var instruction func(opcode, value, pointer int) int
	instruction = func(opcode, value, pointer int) int {
		if opcode == 0 {
			registerA = registerA / int(math.Pow(2, float64(value)))
		} else if opcode == 1 {
			registerB = registerB ^ value
		} else if opcode == 2 {
			registerB = value % 8
		} else if opcode == 3 {
			if registerA != 0 {
				return value
			}
		} else if opcode == 4 {
			registerB = registerB ^ registerC
		} else if opcode == 5 {
			ans = append(ans, strconv.Itoa(value%8))
		} else if opcode == 6 {
			registerB = registerA / int(math.Pow(2, float64(value)))
		} else if opcode == 7 {
			registerC = registerA / int(math.Pow(2, float64(value)))
		}
		return pointer + 2
	}

	programInt := make([]int, len(program))
	for i, str := range program {
		programInt[i], _ = strconv.Atoi(str)
	}

	resultString := strings.Join(ans, ",")
	nextRegisterA := registerA
	for resultString != programStr {
		registerA = nextRegisterA
		registerB = 0
		registerC = 0
		pointer := 0
		for pointer < len(programInt) {
			pointer = instruction(programInt[pointer], getOperand(programInt[pointer+1], programInt[pointer]), pointer)
		}
		resultString = strings.Join(ans, ",")
		ans = ans[:0]
		nextRegisterA *= 8
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.SubmitString(day, 2, resultString)
	}
}
