package Day12

import (
	"AOC2025/Utils"
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 12: Christmas Tree Farm ---
*/

const day = 12

type Shape struct {
	shape [][]int
}

func rotate90(shapeInput Shape) Shape {
	n := len(shapeInput.shape)
	shape := make([][]int, n)
	for i := range shape {
		shape[i] = make([]int, n)
		copy(shape[i], shapeInput.shape[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			shape[i][j], shape[j][i] = shape[j][i], shape[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			shape[i][j], shape[i][n-1-j] = shape[i][n-1-j], shape[i][j]
		}
	}
	return Shape{shape}
}

func generateAllShapes(shapes []Shape) map[int][]Shape {
	allShapes := make(map[int][]Shape)
	for i, shape := range shapes {
		allShapes[i] = append(allShapes[i], shape)
		newShape := shape
		for range 3 {
			newShape = rotate90(newShape)
			allShapes[i] = append(allShapes[i], newShape)
		}
	}
	return allShapes
}

func assignShape(currentStage [][]int, shape Shape, startI, startJ, n, m int) (bool, [][]int) {
	copied := make([][]int, len(currentStage))
	for i := range currentStage {
		copied[i] = make([]int, len(currentStage[i]))
		copy(copied[i], currentStage[i])
	}

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			if startI+i >= n && startJ+j >= m {
				return false, copied
			}
			if shape.shape[i][j] == 1 {
				if copied[i+startI][j+startJ] == 1 {
					return false, copied
				}
				copied[i+startI][j+startJ] = 1
			}
		}
	}
	return true, copied
}

func fillGreedy(shapes map[int][]Shape, n, m int, fillShape []int, currentStage [][]int) bool {
	for index := 0; index < len(shapes); index++ {
		for count := 0; count < fillShape[index]; count++ {
			placed := false
			for i := 0; i < n-2; i++ {
				for j := 0; j < m-2; j++ {
					if currentStage[i][j] == 0 {
						for _, shape := range shapes[index] {
							canAssign, newStage := assignShape(currentStage, shape, i, j, n, m)
							if canAssign {
								currentStage = newStage
								placed = true
								break
							}
						}
						if placed {
							break
						}
					}
				}
				if placed {
					break
				}
			}
			if !placed {
				return false
			}
		}
	}
	return true
}

func solvePuzzle(shapes map[int][]Shape, n, m int, fillShape []int) bool {
	initState := make([][]int, n)
	for i := range initState {
		initState[i] = make([]int, m)
	}
	return fillGreedy(shapes, n, m, fillShape, initState)
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
	shapes := make([]Shape, 6)
	var shape [][]int
	i := 0
	lastLine := 0
	for idx, line := range lines {
		if i > 5 {
			lastLine = idx
			break
		}
		if len(line) == 0 {
			shapes[i] = Shape{shape}
			shape = [][]int{}
			i++
		}
		if len(line) == 3 {
			temp := make([]int, 3)
			for j, char := range line {
				if char == '#' {
					temp[j] = 1
				} else {
					temp[j] = 0
				}
			}
			shape = append(shape, temp)
		}
	}
	ans := 0
	for _, line := range lines[lastLine:] {
		split := strings.Split(line, ": ")
		regionSplit := strings.Split(split[0], "x")
		n, _ := strconv.Atoi(regionSplit[0])
		m, _ := strconv.Atoi(regionSplit[1])
		fillSplit := strings.Split(split[1], " ")
		fillShape := make([]int, len(fillSplit))
		for i, char := range fillSplit {
			value, _ := strconv.Atoi(char)
			fillShape[i] = value
		}
		allShapes := generateAllShapes(shapes)
		if solvePuzzle(allShapes, n, m, fillShape) {
			fmt.Println("CAN SOLVE: ", line)
			ans++
		} else {
			fmt.Println("CANNOT SOLVE: ", line)
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, ans)
	}
}

func Part2(submit bool) {
	ans := "Congratulations! You've solved the puzzle!"
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
}
