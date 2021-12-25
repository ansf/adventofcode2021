package main

import (
	"bufio"
	"os"
)

type Floor [][]int

func NewFloor(dx, dy int) Floor {
	result := make(Floor, dy)
	for y := range result {
		result[y] = make([]int, dx)
	}
	return result
}

func (f Floor) get(x, y int) int {
	return f[y%len(f)][x%len(f[0])]
}

func (f Floor) set(x, y int, v int) {
	f[y%len(f)][x%len(f[0])] = v
}

func (f Floor) Debug() {
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[0]); x++ {
			if f.get(x, y) == 1 {
				print(">")
			} else if f.get(x, y) == 2 {
				print("v")
			} else {
				print(".")
			}
		}
		println()
	}
}

func step(floor Floor) bool {
	didMove := false
	for t := 1; t <= 2; t++ {
		movable := NewFloor(len(floor[0]), len(floor))
		for x := 0; x < len(floor[0]); x++ {
			for y := 0; y < len(floor); y++ {
				if floor.get(x, y) == t {
					if t == 1 && floor.get(x+1, y) == 0 {
						movable.set(x, y, 1)
					} else if t == 2 && floor.get(x, y+1) == 0 {
						movable.set(x, y, 1)
					}
				}
			}
		}

		for x := 0; x < len(floor[0]); x++ {
			for y := 0; y < len(floor); y++ {
				if movable.get(x, y) == 1 {
					didMove = true
					floor.set(x, y, 0)
					if t == 1 {
						floor.set(x+1, y, 1)
					} else if t == 2 {
						floor.set(x, y+1, 2)
					} else {
						panic("illegal state")
					}
				}
			}
		}
	}

	return didMove
}

func main() {
	floor := readInput()

	i := 1
	for step(floor) {
		i++
	}
	println(i)
}

func readInput() Floor {
	result := make([][]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		line := make([]int, 0)
		for _, r := range input {
			if r == '>' {
				line = append(line, 1)
			} else if r == 'v' {
				line = append(line, 2)
			} else {
				line = append(line, 0)
			}
		}
		result = append(result, line)
	}

	return result
}
