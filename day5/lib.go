package day5

import (
	"fmt"
	"io"
	"os"
)

type Point struct {
	x, y int
}

type Line struct {
	Start, End Point
}

func ParseLines(r io.Reader) []Line {
	start := Point{}
	end := Point{}

	result := make([]Line, 0)

	for {
		_, err := fmt.Fscanf(os.Stdin, "%d,%d -> %d,%d", &start.x, &start.y, &end.x, &end.y)
		if err != nil {
			return result
		}
		result = append(result, Line{start, end})
	}
}

func (l Line) IsVerticalOrHorizontal() bool {
	return l.Start.x == l.End.x || l.Start.y == l.End.y
}

func (l Line) GetAllPoints() []Point {
	result := make([]Point, 0)
	dx := l.End.x - l.Start.x
	if dx > 0 {
		dx = 1
	} else if dx < 0 {
		dx = -1
	}
	dy := l.End.y - l.Start.y
	if dy > 0 {
		dy = 1
	} else if dy < 0 {
		dy = -1
	}

	c := l.Start
	result = append(result, c)

	for {
		c = Point{c.x + dx, c.y + dy}
		result = append(result, c)
		if c == l.End {
			break
		}
	}

	return result
}
