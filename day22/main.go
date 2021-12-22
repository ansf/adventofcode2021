package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cuboid struct {
	on bool

	x1, x2 int
	y1, y2 int
	z1, z2 int
}

func (c Cuboid) Contains(x, y, z int) bool {
	return x >= c.x1 && x <= c.x2 &&
		y >= c.y1 && y <= c.y2 &&
		z >= c.z1 && z <= c.z2
}

type Reactor struct {
	min, max int

	data [][][]bool
}

func NewReactor(min, max int) Reactor {
	size := max - min + 1

	result := make([][][]bool, 0)
	for i := 0; i < size; i++ {
		inner := make([][]bool, 0)
		for j := 0; j < size; j++ {
			inner = append(inner, make([]bool, size))
		}

		result = append(result, inner)
	}
	return Reactor{min, max, result}
}

func (r Reactor) Set(x, y, z int, on bool) {
	r.data[x-r.min][y-r.min][z-r.min] = on
}

func (r Reactor) IsOn(x, y, z int) bool {
	return r.data[x-r.min][y-r.min][z-r.min]
}

func main() {
	cuboids := readInput()
	reactor := NewReactor(-50, 50)
	for _, c := range cuboids {
		for x := -50; x <= 50; x++ {
			for y := -50; y <= 50; y++ {
				for z := -50; z <= 50; z++ {
					if c.Contains(x, y, z) {
						reactor.Set(x, y, z, c.on)
					}
				}
			}
		}
	}

	count := 0
	for x := -50; x <= 50; x++ {
		for y := -50; y <= 50; y++ {
			for z := -50; z <= 50; z++ {
				if reactor.IsOn(x, y, z) {
					count++
				}
			}
		}
	}
	println(count)
}

func readInput() []Cuboid {
	result := make([]Cuboid, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var onOff string
		var x1, x2, y1, y2, z1, z2 int
		_, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &onOff, &x1, &x2, &y1, &y2, &z1, &z2)
		if err != nil {
			panic("failed to parse")
		}

		result = append(result, Cuboid{onOff == "on", x1, x2, y1, y2, z1, z2})
	}
	return result
}
