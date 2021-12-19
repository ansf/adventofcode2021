package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y, z int
}

func (p1 Point) Sub(p2 Point) Point {
	return Point{
		p1.x - p2.x,
		p1.y - p2.y,
		p1.z - p2.z,
	}
}

func (p Point) ManhattenLength() int {
	return Abs(p.x) + Abs(p.y) + Abs(p.z)
}

func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

// M is a 4x4 matrix
type M []int

// V is the value of the matric at position x,y
func (m M) V(x, y int) int {
	return m[x+y*4]
}

func (m M) SetV(x, y int, v int) {
	m[x+y*4] = v
}

func (m M) MulP(p Point) Point {
	v := []int{p.x, p.y, p.z, 1}
	result := []int{0, 0, 0, 0}
	for y := 0; y < 4; y++ {
		result[y] = v[0]*m.V(0, y) + v[1]*m.V(1, y) + v[2]*m.V(2, y) + v[3]*m.V(3, y)
	}
	return Point{result[0], result[1], result[2]}
}

func (ma M) MulM(mb M) M {
	m := make(M, 16)

	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			v := 0
			for i := 0; i < 4; i++ {
				v += ma.V(i, y) * mb.V(x, i)
			}
			m.SetV(x, y, v)
		}
	}

	return m
}

// ID creates an identity matric
func ID() M {
	return T(0, 0, 0)
}

// T creates a translation matrix
func T(x, y, z int) M {
	return M{
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1,
	}
}

// R creates a rotation matrix where x, y and z are in multiples of 90 degree
func R(x, y, z int) M {
	a, b, c := float64(x)*math.Pi/2, float64(y)*math.Pi/2, float64(z)*math.Pi/2

	ra := M{
		int(math.Cos(a)), -int(math.Sin(a)), 0, 0,
		int(math.Sin(a)), int(math.Cos(a)), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	rb := M{
		int(math.Cos(b)), 0, int(math.Sin(b)), 0,
		0, 1, 0, 0,
		-int(math.Sin(b)), 0, int(math.Cos(b)), 0,
		0, 0, 0, 1,
	}

	rc := M{
		1, 0, 0, 0,
		0, int(math.Cos(c)), -int(math.Sin(c)), 0,
		0, int(math.Sin(c)), int(math.Cos(c)), 0,
		0, 0, 0, 1,
	}

	return ra.MulM(rb).MulM(rc)
}

func Rotations() []M {
	return []M{
		R(0, 0, 0), R(0, 0, 1), R(0, 0, 2), R(0, 0, 3),
		R(0, 1, 0), R(0, 1, 1), R(0, 1, 2), R(0, 1, 3),
		R(0, 2, 0), R(0, 2, 1), R(0, 2, 2), R(0, 2, 3),
		R(0, 3, 0), R(0, 3, 1), R(0, 3, 2), R(0, 3, 3),
		R(1, 0, 0), R(1, 0, 1), R(1, 0, 2), R(1, 0, 3),
		R(3, 0, 0), R(3, 0, 1), R(3, 0, 2), R(3, 0, 3),
	}
}

func findTransform(scanner1, scanner2 []Point) *M {
	scanner1Lookup := make(map[Point]bool)
	for _, p := range scanner1 {
		scanner1Lookup[p] = true
	}

	// try all rotations
	for _, rotation := range Rotations() {

		// bring all beacons in the given rotation
		rotated := make([]Point, 0)
		for _, beacon := range scanner2 {
			rotated = append(rotated, rotation.MulP(beacon))
		}

		// build list of possible translations by aligning one point at a time
		translations := make([]Point, 0)
		for _, p1 := range scanner1 {
			for _, p2 := range rotated {
				translations = append(translations, p1.Sub(p2))
			}
		}

		// try all possible translations to overlap the beacons
		for _, t := range translations {
			transform := T(t.x, t.y, t.z)
			transformed := make([]Point, 0)
			for _, beacon := range rotated {
				transformed = append(transformed, transform.MulP(beacon))
			}

			overlap := make([]Point, 0)
			count := 0
			for _, p := range transformed {
				if scanner1Lookup[p] {
					count++
					overlap = append(overlap, p)
				}
			}

			// 12 or more overlaps mean we found the transformation of scanner2 relative to scanner1
			if count >= 12 {
				m := transform.MulM(rotation)
				return &m
			}
		}
	}

	return nil
}

func main() {
	scanners := readInput()

	// known scanner indexes and their transform
	known := make(map[int]M)

	// we use scanner0 as our reference system
	known[0] = ID()

	for {
		for unknownIdx := range scanners {
			_, ok := known[unknownIdx]
			if ok {
				continue
			}

			for refIdx := range known {
				m := findTransform(scanners[refIdx], scanners[unknownIdx])
				if m == nil {
					continue
				}
				for i := range scanners[unknownIdx] {
					scanners[unknownIdx][i] = m.MulP(scanners[unknownIdx][i])
				}
				known[unknownIdx] = *m
				break
			}
		}
		if len(known) == len(scanners) {
			break
		}
	}

	// now all beacons are in the same system, just count them
	lookup := make(map[Point]struct{})
	for _, beacons := range scanners {
		for _, beacon := range beacons {
			lookup[beacon] = struct{}{}
		}
	}
	println("Beacons:", len(lookup))

	// find max distance of scanners
	maxDistance := 0
	for i := 0; i < len(scanners); i++ {
		for j := 0; j < len(scanners); j++ {
			if i == j {
				continue
			}

			// build vectors of each system to the reference system
			vToRefI := known[i].MulP(Point{0, 0, 0})
			vToRefJ := known[j].MulP(Point{0, 0, 0})

			// then simple vector math
			manhattenLength := vToRefI.Sub(vToRefJ).ManhattenLength()
			if manhattenLength > maxDistance {
				maxDistance = manhattenLength
			}
		}
	}
	println("Max scanner distance:", maxDistance)
}

func readInput() [][]Point {
	result := make([][]Point, 0)

	current := make([]Point, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "---") {
			continue
		}

		if line == "" {
			result = append(result, current)
			current = make([]Point, 0)
			continue
		}

		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic("failed to parse input")
		}

		current = append(current, Point{x, y, z})
	}

	result = append(result, current)

	return result
}
