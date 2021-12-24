package main

type Register struct {
	x, y, z, w int
}

// exec is a manual decompilation of the 18 instructions that transform an input.
// these 18 instruction are identical for all 14 inputs, except for the parameters a, b and c.
func exec(idx int, input int, a, b, c int, r Register) Register {
	r.x = r.z%26 + b

	if r.x != input {
		r.x = 1
		r.z /= a
		r.z *= 26
		r.z = r.z + (input + c)
	} else {
		r.x = 0
		r.z = r.z / a
	}

	return r
}

var (
	as = []int{1, 1, 1, 1, 26, 26, 26, 1, 1, 26, 26, 26, 1, 26}
	bs = []int{13, 15, 15, 11, -16, -11, -6, 11, 10, -10, -8, -11, 12, -15}
	cs = []int{5, 14, 15, 16, 8, 9, 2, 13, 16, 6, 6, 9, 11, 5}
)

var searchRange = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

func init() {
	if false {
		searchRange = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	}
}

func main() {
	for _, input := range searchRange {
		if solve(0, input, Register{}, []int{}) {
			return
		}
	}

	println("no valid number found")
}

func solve(idx int, input int, state Register, solution []int) bool {
	if idx == 14 {
		for _, s := range solution {
			print(s)
		}
		println()
		return true
	}

	a := as[idx]
	b := bs[idx]
	c := cs[idx]

	next := exec(idx, input, a, b, c, state)

	// by inspecting register x, we can check if we hit the division
	// to keep our z from growing too high
	if next.x == 1 && a == 26 {
		return false
	}

	for _, nextInput := range searchRange {
		if solve(idx+1, nextInput, next, append(solution, input)) {
			return true
		}
	}
	return false
}
