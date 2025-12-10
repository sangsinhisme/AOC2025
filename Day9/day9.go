package Day9

import (
	"AOC2025/Utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 9: Movie Theater ---
*/
const day = 9

type Point struct {
	X, Y int
}

func findSquare(p1, p2 Point) int {
	return (Utils.Abs(p1.X, p2.X) + 1) * (Utils.Abs(p1.Y, p2.Y) + 1)
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
	Points := make([]Point, n)
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		Points[i] = Point{x, y}
	}
	sort.Slice(Points, func(i, j int) bool {
		if Points[i].X == Points[j].X {
			return Points[i].Y < Points[j].Y
		}
		return Points[i].X < Points[j].X
	})
	ans := 0
	for i := 0; i < len(Points)-1; i++ {
		for j := i + 1; j < len(Points); j++ {
			ans = max(ans, findSquare(Points[i], Points[j]))
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
	Points := make(map[Point]bool)
	m, n := 0, 0

	for _, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		m = max(m, x+1)
		n = max(n, y+1)
		Points[Point{x, y}] = true
	}

	directs := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	maxTop := make([]int, m)
	minBottem := make([]int, m)
	for i := range minBottem {
		minBottem[i] = math.MaxInt
	}
	minLeft := make([]int, n)
	for i := range minLeft {
		minLeft[i] = math.MaxInt
	}
	maxRight := make([]int, n)
	var moveUntilCatch func(p Point)
	moveUntilCatch = func(p Point) {
		var nextMoves [][2]int
		for _, direct := range directs {
			initX, initY := p.X+direct[0], p.Y+direct[1]
			for initX >= 0 && initY >= 0 && initX < m && initY < n {
				if _, ok := Points[Point{initX, initY}]; ok {
					nextMoves = append(nextMoves, direct)
					break
				}
				initX, initY = initX+direct[0], initY+direct[1]
			}
		}
		for _, direct := range nextMoves {
			initX, initY := p.X, p.Y
			for initX >= 0 && initY >= 0 && initX < m && initY < n {
				maxTop[initX] = max(maxTop[initX], initY)
				minBottem[initX] = min(minBottem[initX], initY)
				minLeft[initY] = min(minLeft[initY], initX)
				maxRight[initY] = max(maxRight[initY], initX)
				initX, initY = initX+direct[0], initY+direct[1]
				if _, ok := Points[Point{initX, initY}]; ok {
					break
				}
			}
		}
	}
	for p := range Points {
		moveUntilCatch(p)
	}

	pointArr := make([]Point, 0, len(Points))
	for point := range Points {
		pointArr = append(pointArr, point)
	}
	sort.Slice(pointArr, func(i, j int) bool {
		if pointArr[i].X == pointArr[j].X {
			return pointArr[i].Y < pointArr[j].Y
		}
		return pointArr[i].X < pointArr[j].X
	})
	ans := 0
	fmt.Println(maxTop, minBottem, minLeft, maxRight)
	for i := range pointArr {
		for j := i + 1; j < len(pointArr); j++ {
			minX := min(pointArr[i].X, pointArr[j].X)
			maxX := max(pointArr[i].X, pointArr[j].X)
			minY := min(pointArr[i].Y, pointArr[j].Y)
			maxY := max(pointArr[i].Y, pointArr[j].Y)

			if minX >= minLeft[minY] &&
				minY >= minBottem[minX] &&
				maxX <= maxRight[maxY] &&
				maxY <= maxTop[maxX] &&
				minX >= minLeft[maxY] &&
				minY >= minBottem[maxX] &&
				maxX <= maxRight[minY] &&
				maxY <= maxTop[minX] {
				ans = max(ans, findSquare(pointArr[i], pointArr[j]))
			}
		}
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
