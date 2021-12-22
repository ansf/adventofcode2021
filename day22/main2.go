package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Cuboid struct {
	on bool

	x1, x2 int
	y1, y2 int
	z1, z2 int
}

func (c Cuboid) Contains(x, y, z int) bool {
	return x >= c.x1 && x < c.x2 &&
		y >= c.y1 && y < c.y2 &&
		z >= c.z1 && z < c.z2
}

// find finds the cuboid that is controlling position x,y,z
func find(cuboids []Cuboid, x, y, z int) *Cuboid {
	for i := len(cuboids) - 1; i >= 0; i-- {
		if cuboids[i].Contains(x, y, z) {
			return &cuboids[i]
		}
	}

	return nil
}

func main() {
	cuboids := readInput()

	// generate ordered lists of intersection points
	// basically we are making a grid a smaller cuboids such
	// that the bigger cuboid confiuration can be divided
	// into these smaller cuboids without overlap
	uniqueXs := make(map[int]bool, 0)
	uniqueYs := make(map[int]bool, 0)
	uniqueZs := make(map[int]bool, 0)
	for _, c := range cuboids {
		uniqueXs[c.x1] = true
		uniqueXs[c.x2] = true
		uniqueYs[c.y1] = true
		uniqueYs[c.y2] = true
		uniqueZs[c.z1] = true
		uniqueZs[c.z2] = true
	}

	xs := make([]int, 0)
	for x := range uniqueXs {
		xs = append(xs, x)
	}
	sort.Ints(xs)

	ys := make([]int, 0)
	for y := range uniqueYs {
		ys = append(ys, y)
	}
	sort.Ints(ys)

	zs := make([]int, 0)
	for z := range uniqueZs {
		zs = append(zs, z)
	}
	sort.Ints(zs)

	// this is still ~800^3 iterations for the real input, which finishes in about 5 minutes
	// we might want to figure out something more smart...
	count := 0
	for x := 0; x < len(xs)-1; x++ {
		println(x)
		for y := 0; y < len(ys)-1; y++ {
			for z := 0; z < len(zs)-1; z++ {
				cuboid := find(cuboids, xs[x], ys[y], zs[z])
				if cuboid != nil && cuboid.on {
					toX := xs[x+1]
					toY := ys[y+1]
					toZ := zs[z+1]

					count += (toX - xs[x]) * (toY - ys[y]) * (toZ - zs[z])
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

		// we use end-exclusive coordinates
		result = append(result, Cuboid{onOff == "on", x1, x2 + 1, y1, y2 + 1, z1, z2 + 1})
	}
	return result
}
