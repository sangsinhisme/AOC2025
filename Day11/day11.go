package Day11

import (
	"AOC2025/Utils"
	"fmt"
	"strings"
)

/*
--- Day 11: Reactor ---
*/

const day = 11

func bfs(edges map[string][]string, start string) int {
	reachOut := 0
	queue := [][]string{{start, start}}
	for len(queue) > 0 {
		node := queue[0]
		for _, neighbor := range edges[node[0]] {
			if neighbor == "out" {
				reachOut++
				continue
			}
			nextPath := node[1] + "," + neighbor
			if !strings.Contains(node[1], ","+neighbor) {
				queue = append(queue, []string{neighbor, nextPath})
			}
		}
		queue = queue[1:]
	}
	return reachOut
}

func countPaths(edges map[string][]string, start, end string) int {
	memo := make(map[string]int)
	return countDP(edges, start, end, memo)
}

func countDP(edges map[string][]string, current, end string, memo map[string]int) int {
	if current == end {
		return 1
	}
	if count, found := memo[current]; found {
		return count
	}
	count := 0
	for _, neighbor := range edges[current] {
		count += countDP(edges, neighbor, end, memo)
	}
	memo[current] = count
	return count
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
	edges := make(map[string][]string)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		node := split[0]
		for _, neighbor := range strings.Split(split[1], " ") {
			edges[node] = append(edges[node], neighbor)
		}
	}
	ans := bfs(edges, "you")
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
	edges := make(map[string][]string)
	nodes := make(map[string]bool)
	for _, line := range lines {
		split := strings.Split(line, ": ")
		node := split[0]
		nodes[node] = true
		for _, neighbor := range strings.Split(split[1], " ") {
			edges[node] = append(edges[node], neighbor)
		}
	}

	// let A is svr
	// let B is fft
	// let C is dac
	// let D is out
	ATOB := countPaths(edges, "svr", "fft")
	BTOC := countPaths(edges, "fft", "dac")
	CTOD := countPaths(edges, "dac", "out")

	ans := ATOB * BTOC * CTOD
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
