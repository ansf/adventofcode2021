package main

import (
	"os"

	"github.com/ansf/adventofcode_2021/lib"
)

func main() {
	positions, err := lib.ReadListOfNumbers(os.Stdin)
	if err != nil {
		panic("could not read input")
	}

	sum := 0
	for _, p := range positions {
		sum += p
	}
	avg := sum / len(positions)

	fuel := RequiredFuel(positions, avg)
	for i := 0; ; i++ {
		nextFuel := RequiredFuel(positions, avg+i)
		if nextFuel > fuel {
			break
		}
		fuel = nextFuel
	}
	for i := 0; ; i-- {
		nextFuel := RequiredFuel(positions, avg+i)
		if nextFuel > fuel {
			break
		}
		fuel = nextFuel
	}
	println(fuel)
}

func RequiredFuel(positions []int, target int) int {
	fuel := 0
	for _, p := range positions {
		fuel += Abs(p - target)
	}
	return fuel
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
