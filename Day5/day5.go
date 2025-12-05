package Day5

import (
	"AOC2025/Utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 5: Cafeteria ---
*/
const day = 5

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
	var points [][2]int
	var ingredients []int
	breakPoint := false
	for _, line := range lines {
		if len(line) == 0 {
			breakPoint = true
			continue
		}
		if !breakPoint {
			split := strings.Split(line, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])
			points = append(points, [2]int{start, end})
		} else {
			ingredient, _ := strconv.Atoi(line)
			ingredients = append(ingredients, ingredient)
		}
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][1] == points[j][1] {
			return points[i][0] > points[j][0]
		}
		return points[i][1] < points[j][1]
	})
	sort.Slice(ingredients, func(i, j int) bool {
		return ingredients[i] < ingredients[j]
	})
	ans := 0
	pointLeftIdx, maxIdx := 0, len(points)
	visited := make(map[int]bool)
	for _, ingredient := range ingredients {
		nextPointLeft := pointLeftIdx
		for nextPointLeft < maxIdx {
			if ingredient > points[nextPointLeft][1] {
				nextPointLeft++
			} else {
				if ingredient >= points[nextPointLeft][0] {
					fmt.Printf("Ingredient ID %d is fresh because it falls into range %d-%d", ingredient, points[nextPointLeft][0], points[nextPointLeft][1])
					visited[ingredient] = true
					ans += 1
				} else {
					fmt.Printf("Ingredient ID %d is spoiled because it does not fall into any range.", ingredient)
				}
				fmt.Println()
				break
			}
		}
		pointLeftIdx = nextPointLeft
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	fmt.Println("Lest start right check")
	pointRightIdx := 0
	for _, ingredient := range ingredients {
		nextPointRight := pointRightIdx
		if visited[ingredient] || ingredient < points[nextPointRight][0] {
			continue
		}
		for nextPointRight < maxIdx {
			if points[nextPointRight][1] < ingredient {
				nextPointRight++
			} else {
				if ingredient >= points[nextPointRight][0] {
					fmt.Printf("Ingredient ID %d is fresh because it falls into range %d-%d", ingredient, points[nextPointRight][0], points[nextPointRight][1])
					ans += 1
				} else {
					fmt.Printf("Ingredient ID %d is spoiled because it does not fall into any range.", ingredient)
				}
				fmt.Println()
				break
			}
		}
		pointRightIdx = nextPointRight
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
	var points [][2]int
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		split := strings.Split(line, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		points = append(points, [2]int{start, end})
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][0] > points[j][0]
		}
		return points[i][0] < points[j][0]
	})
	ans := 0
	maxRight := 0
	for _, point := range points {
		if maxRight > point[1] {
			continue
		}
		fmt.Printf("Add range %d-%d\n", max(maxRight, point[0]), point[1])
		ans += point[1] - max(maxRight, point[0]) + 1
		maxRight = point[1] + 1
	}
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
