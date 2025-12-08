package Day8

import (
	"AOC2025/Utils"
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

/*
--- Day 8: Playground ---
*/
const day = 8

type Point struct {
	X, Y, Z int
}

type PairDistance struct {
	DistanceSq float64
	P1, P2     Point
}

const K = 1000

type MaxHeap []PairDistance

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].DistanceSq > h[j].DistanceSq } // Max-heap logic
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(PairDistance))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
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
	boxes := make([]Point, n)
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		boxes[i] = Point{x, y, z}
	}
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := boxes[i].X - boxes[j].X
			dy := boxes[i].Y - boxes[j].Y
			dz := boxes[i].Z - boxes[j].Z

			distanceSq := float64(dx*dx + dy*dy + dz*dz)
			newPair := PairDistance{
				DistanceSq: distanceSq,
				P1:         boxes[i],
				P2:         boxes[j],
			}
			if h.Len() < K {
				heap.Push(h, newPair)
			} else if newPair.DistanceSq < (*h)[0].DistanceSq {
				heap.Pop(h) //pop largest element
				heap.Push(h, newPair)
			}
		}
	}
	nodes := make(map[Point]bool)
	edges := make(map[Point][]Point)
	for i := h.Len() - 1; i >= 0; i-- {
		edge := heap.Pop(h).(PairDistance)
		nodes[edge.P1] = true
		nodes[edge.P2] = true
		edges[edge.P1] = append(edges[edge.P1], edge.P2)
		edges[edge.P2] = append(edges[edge.P2], edge.P1)
	}

	visited := make(map[Point]bool)
	var dfs func(p Point) int
	dfs = func(p Point) int {
		countVisit := len(visited)
		visited[p] = true
		for _, neighbor := range edges[p] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		return len(visited) - countVisit
	}

	for node := range nodes {
		if !visited[node] {
			circles := dfs(node)
			heap.Push(h, PairDistance{DistanceSq: float64(circles), P1: node, P2: node})
		}
	}

	ans := 1
	for i := 0; i <= 2; i++ {
		bestCircle := heap.Pop(h).(PairDistance)
		fmt.Println(bestCircle.P1, bestCircle.P2, bestCircle.DistanceSq)
		ans *= int(bestCircle.DistanceSq)
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
	n := len(lines)
	boxes := make([]Point, n)
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		boxes[i] = Point{x, y, z}
	}
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := boxes[i].X - boxes[j].X
			dy := boxes[i].Y - boxes[j].Y
			dz := boxes[i].Z - boxes[j].Z

			distanceSq := float64(dx*dx + dy*dy + dz*dz)
			newPair := PairDistance{
				DistanceSq: distanceSq,
				P1:         boxes[i],
				P2:         boxes[j],
			}
			heap.Push(h, newPair)
		}
	}
	edgesSorted := make([]PairDistance, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		edgesSorted[i] = heap.Pop(h).(PairDistance)
	}

	var helper func(numEdge int) bool
	helper = func(numEdge int) bool {
		visited := make(map[Point]bool)
		edges := make(map[Point][]Point)
		for i := range numEdge {
			edge := edgesSorted[i]
			edges[edge.P1] = append(edges[edge.P1], edge.P2)
			edges[edge.P2] = append(edges[edge.P2], edge.P1)
		}
		var dfs func(p Point)
		dfs = func(p Point) {
			visited[p] = true
			for _, neighbor := range edges[p] {
				if !visited[neighbor] {
					dfs(neighbor)
				}
			}
		}
		dfs(edgesSorted[0].P1)
		return len(visited) == n
	}
	left, right := 0, len(edgesSorted)
	for left < right {
		mid := (left + right) / 2
		if helper(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	lastEdge := edgesSorted[left-1]
	ans := lastEdge.P1.X * lastEdge.P2.X
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, ans)
	}
}
