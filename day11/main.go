package main

import (
	"bufio"
	"os"
)

type O struct {
	energy  int
	flashed bool
}

func NewO(energy int) O {
	return O{energy, false}
}

type Point struct {
	x, y int
}

type Grid struct {
	data [][]O
}

func (g *Grid) Print() {
	for y := 0; y < g.Dy(); y++ {
		for x := 0; x < g.Dx(); x++ {
			print(g.data[y][x].energy)
		}
		println()
	}
	println()
}

func (g *Grid) Dx() int {
	return len(g.data[0])
}

func (g *Grid) Dy() int {
	return len(g.data)
}

func (g *Grid) GetAdjacents(p Point) []Point {
	result := make([]Point, 0)
	if p.x > 0 {
		result = append(result, Point{p.x - 1, p.y})
		if p.y > 0 {
			result = append(result, Point{p.x - 1, p.y - 1})
		}
		if p.y < g.Dy()-1 {
			result = append(result, Point{p.x - 1, p.y + 1})
		}
	}
	if p.x < g.Dx()-1 {
		result = append(result, Point{p.x + 1, p.y})
		if p.y > 0 {
			result = append(result, Point{p.x + 1, p.y - 1})
		}
		if p.y < g.Dy()-1 {
			result = append(result, Point{p.x + 1, p.y + 1})
		}
	}

	if p.y > 0 {
		result = append(result, Point{p.x, p.y - 1})
	}
	if p.y < g.Dy()-1 {
		result = append(result, Point{p.x, p.y + 1})
	}

	return result
}

func (g *Grid) Get(x, y int) *O {
	return &g.data[y][x]
}

func (g *Grid) ForEarch(f func(x, y int, o *O)) {
	for x := 0; x < g.Dx(); x++ {
		for y := 0; y < g.Dy(); y++ {
			f(x, y, g.Get(x, y))
		}
	}
}

func (g *Grid) Step() int {
	g.ForEarch(func(_, _ int, o *O) { o.energy++ })

	flashes := 0
	for {
		flashed := false
		g.ForEarch(func(x, y int, o *O) {
			if !o.flashed && o.energy > 9 {
				flashes++
				o.flashed = true
				for _, adjacent := range g.GetAdjacents(Point{x, y}) {
					g.Get(adjacent.x, adjacent.y).energy++
				}
				flashed = true
			}
		})
		if !flashed {
			break
		}
	}

	g.ForEarch(func(x, y int, o *O) {
		if o.flashed {
			o.energy = 0
			o.flashed = false
		}
	})

	return flashes
}

func main() {
	g := readInput()
	g.Print()

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += g.Step()
		g.Print()
	}
	println(flashes)
}

func readInput() Grid {
	grid := Grid{make([][]O, 0)}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		line := make([]O, 0)
		for _, r := range input {
			line = append(line, NewO(int(r-'0')))
		}
		grid.data = append(grid.data, line)
		line = make([]O, 0)
	}

	return grid
}
