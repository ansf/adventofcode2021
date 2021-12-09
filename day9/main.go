package main

import (
	"bufio"
	"os"
)

type FloorPlan [][]int

func (f FloorPlan) GetAdjacents(x, y int) []int {
	adjacents := make([]int, 0)
	if x > 0 {
		adjacents = append(adjacents, f[y][x-1])
	}
	if x < len(f[0])-1 {
		adjacents = append(adjacents, f[y][x+1])
	}
	if y > 0 {
		adjacents = append(adjacents, f[y-1][x])
	}
	if y < len(f)-1 {
		adjacents = append(adjacents, f[y+1][x])
	}
	return adjacents
}

func (f FloorPlan) IsLowPoint(x, y int) bool {
	adjacents := f.GetAdjacents(x, y)
	for _, a := range adjacents {
		if f[y][x] >= a {
			return false
		}
	}
	return true
}

func main() {
	floorPlan := readFloorPlan()

	sum := 0
	for x := 0; x < len(floorPlan[0]); x++ {
		for y := 0; y < len(floorPlan); y++ {
			if floorPlan.IsLowPoint(x, y) {
				sum += floorPlan[y][x] + 1
			}
		}
	}

	println(sum)
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
