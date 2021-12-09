package main

import (
	"bufio"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

func (p Point) GetAdjacents() []Point {
	return []Point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

type VisitedMap [][]bool

func NewVisitedMap(size Point) VisitedMap {
	visited := make([][]bool, size.y)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, size.x)
	}
	return visited
}

func (v VisitedMap) Visit(p Point) {
	if p.x < 0 || p.x >= len(v[0]) || p.y < 0 || p.y >= len(v) {
		return
	}
	v[p.y][p.x] = true
}

func (v VisitedMap) IsVisited(p Point) bool {
	if p.x < 0 || p.x >= len(v[0]) || p.y < 0 || p.y >= len(v) {
		return true
	}
	return v[p.y][p.x]
}

type FloorPlan [][]int

func (f FloorPlan) GetHeight(p Point) int {
	if p.x < 0 || p.x >= len(f[0]) {
		return 9
	}
	if p.y < 0 || p.y >= len(f) {
		return 9
	}
	return f[p.y][p.x]
}

func (f FloorPlan) IsLowPoint(x, y int) bool {
	adjacents := Point{x, y}.GetAdjacents()
	for _, a := range adjacents {
		if f.GetHeight(Point{x, y}) >= f.GetHeight(a) {
			return false
		}
	}
	return true
}

func (f FloorPlan) FindBasinSize(p Point) int {
	size := 0

	points := []Point{p}
	nextPoints := make([]Point, 0)

	visited := NewVisitedMap(Point{len(f[0]), len(f)})
	visited.Visit(points[0])

	for {
		for _, p := range points {
			if f.GetHeight(p) < 9 {
				size++
			}

			adjacents := p.GetAdjacents()
			for _, a := range adjacents {
				if f.GetHeight(a) == 9 {
					continue
				}
				if visited.IsVisited(a) {
					continue
				}
				visited.Visit(a)
				nextPoints = append(nextPoints, a)
			}
		}
		if len(nextPoints) == 0 {
			break
		}
		points = nextPoints
		nextPoints = make([]Point, 0)
	}

	return size
}

func main() {
	floorPlan := readFloorPlan()

	lowPoints := make([]Point, 0)
	for x := 0; x < len(floorPlan[0]); x++ {
		for y := 0; y < len(floorPlan); y++ {
			if floorPlan.IsLowPoint(x, y) {
				lowPoints = append(lowPoints, Point{x, y})
			}
		}
	}

	sizes := make([]int, 0)
	for _, p := range lowPoints {
		sizes = append(sizes, floorPlan.FindBasinSize(p))
	}

	sort.Ints(sizes)

	result := 1
	for i := len(sizes) - 3; i < len(sizes); i++ {
		result *= sizes[i]
	}
	println(result)
}

func readFloorPlan() FloorPlan {
	result := make([][]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		line := make([]int, 0)
		for _, s := range input {
			n := int(s - '0')
			line = append(line, n)
		}

		result = append(result, line)
	}

	return result
}
