package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Image struct {
	pixels map[Point]bool

	min Point
	max Point

	fault bool
}

func NewImage() Image {
	pixels := make(map[Point]bool)
	return Image{pixels, Point{math.MaxInt, math.MaxInt}, Point{math.MinInt, math.MinInt}, false}
}

func (im Image) IsOn(p Point) bool {
	if p.x < im.min.x || p.x > im.max.x || p.y < im.min.y || p.y > im.max.y {
		return im.fault
	}
	return im.pixels[p]
}

func (im *Image) On(p Point) {
	im.pixels[p] = true
	if p.x > im.max.x {
		im.max.x = p.x
	}
	if p.x < im.min.x {
		im.min.x = p.x
	}
	if p.y > im.max.y {
		im.max.y = p.y
	}
	if p.y < im.min.y {
		im.min.y = p.y
	}
}

func (im Image) Enhance(algorithm string) Image {
	result := NewImage()
	for x := im.min.x - 3; x <= im.max.x+3; x++ {
		for y := im.min.y - 3; y <= im.max.y+3; y++ {
			point := Point{x, y}
			pixel := algorithm[im.Decode(point)]
			if pixel == '#' {
				result.On(point)
				// if we are on the border, set the faulting pixel
				if x == im.min.x-3 && y == im.min.y-3 {
					result.fault = true
				}
			}
		}
	}

	return result
}

func (im Image) Decode(point Point) int {
	result := 0
	for _, p := range point.Adjacents() {
		result <<= 1
		if im.IsOn(p) {
			result |= 1
		}
	}
	return result
}

func (im Image) Debug() {
	for y := im.min.y - 3; y <= im.max.y+3; y++ {
		for x := im.min.x - 3; x <= im.max.x+3; x++ {
			if im.IsOn(Point{x, y}) {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
	println()
}

type Point struct {
	x, y int
}

func (p Point) Adjacents() []Point {
	return []Point{
		{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y - 1},

		{p.x - 1, p.y},
		{p.x, p.y},
		{p.x + 1, p.y},

		{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		{p.x + 1, p.y + 1},
	}
}

func main() {
	algorithm, image := readInput()

	for i := 1; i <= 50; i++ {
		image = image.Enhance(algorithm)
		if i == 2 || i == 50 {
			fmt.Printf("Lights after %d enhancements: %d\n", i, len(image.pixels))
		}
	}
}

func readInput() (string, Image) {
	var algorithm string
	fmt.Fscanf(os.Stdin, "%s", &algorithm)

	image := NewImage()
	y := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				image.On(Point{x, y})
			}
		}

		y++
	}

	return algorithm, image
}
