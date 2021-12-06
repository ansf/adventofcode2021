package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Simulation struct {
	day int
	// fishes at index i breed after `(days-1) % 7` days
	fishes []int
	// younglings start at index 0, move to index 1 then become fishes
	younglings []int
}

func NewSimulation() Simulation {
	return Simulation{
		0,
		make([]int, 7),
		make([]int, 2),
	}
}

func (s *Simulation) AddFish(fish int) {
	s.fishes[fish]++
}

func (s *Simulation) Step() {
	s.day++

	breedingSlot := (s.day - 1) % 7
	newBorn := s.fishes[breedingSlot]
	s.fishes[(s.day-1)%7] += s.younglings[1]
	s.younglings[1] = s.younglings[0]
	s.younglings[0] = newBorn
}

func (s *Simulation) Size() int {
	count := 0
	for _, c := range s.fishes {
		count += c
	}
	return count + s.younglings[0] + s.younglings[1]
}

func main() {
	input := ""
	_, err := fmt.Fscanf(os.Stdin, "%s", &input)
	if err != nil {
		panic("could not read input")
	}

	sim := NewSimulation()

	for _, v := range strings.Split(input, ",") {
		fish, err := strconv.Atoi(v)
		if err != nil {
			panic("could not parse fishes")
		}
		sim.AddFish(fish)
	}

	for i := 0; i < 256; i++ {
		sim.Step()
	}

	println(sim.Size())
}
