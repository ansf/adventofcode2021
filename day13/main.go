package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func (p Point) Fold(f Fold) Point {
	if f.axis == 1 {
		if p.y < f.along {
			return p
		} else {
			d := p.y - f.along
			return Point{p.x, f.along - d}
		}
	} else {
		if p.x < f.along {
			return p
		} else {
			d := p.x - f.along
			return Point{f.along - d, p.y}
		}
	}

}

type Fold struct {
	along int
	axis  int
}

func FoldPoints(points []Point, fold Fold) []Point {
	foldedPoints := make([]Point, 0)
	for _, p := range points {
		foldedPoints = append(foldedPoints, p.Fold(fold))
	}
	return foldedPoints
}

func HasPoint(points []Point, p Point) bool {
	for _, point := range points {
		if p == point {
			return true
		}
	}
	return false
}

func PrintPoints(points []Point) {
	for y := 0; y < 15; y++ {
		for x := 0; x < 30; x++ {
			if HasPoint(points, Point{x, y}) {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func main() {
	points, folds := readInput()

	foldedPoints := FoldPoints(points, folds[0])

	countMap := make(map[Point]struct{})
	for _, p := range foldedPoints {
		countMap[p] = struct{}{}
	}

	println(len(countMap))
}

func readInput() ([]Point, []Fold) {
	scanner := bufio.NewScanner(os.Stdin)

	points := make([]Point, 0)
	folds := make([]Fold, 0)

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			continue
		} else if strings.HasPrefix(input, "fold") {
			var axis string
			var along int
			fmt.Sscanf(input, "fold along %1s=%d", &axis, &along)
			if axis == "x" {
				folds = append(folds, Fold{along, 0})
			} else {
				folds = append(folds, Fold{along, 1})
			}
		} else {
			var p Point
			fmt.Sscanf(input, "%d,%d", &p.x, &p.y)
			points = append(points, p)
		}
	}

	return points, folds
}
