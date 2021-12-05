package main

import (
	"fmt"
	"os"

	"github.com/ansf/adventofcode2021/day5"
)

func main() {
	lines := day5.ParseLines(os.Stdin)

	mapData := make(map[day5.Point]int)

	for _, line := range lines {
		points := line.GetAllPoints()
		for _, p := range points {
			mapData[p]++
		}
	}

	count := 0
	for _, v := range mapData {
		if v >= 2 {
			count++
		}
	}

	fmt.Printf("%d\n", count)

}
