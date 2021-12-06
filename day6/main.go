package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ""
	_, err := fmt.Fscanf(os.Stdin, "%s", &input)
	if err != nil {
		panic("could not read input")
	}

	fishes := make([]int, 0)
	for _, v := range strings.Split(input, ",") {
		fish, err := strconv.Atoi(v)
		if err != nil {
			panic("could not parse fishes")
		}
		fishes = append(fishes, fish)
	}

	for i := 0; i < 80; i++ {
		fishes = stepDay(fishes)
	}

	println(len(fishes))
}

func stepDay(fishes []int) []int {
	newFishes := make([]int, 0)

	for i := 0; i < len(fishes); i++ {
		if fishes[i] == 0 {
			newFishes = append(newFishes, 8)
			fishes[i] = 6
		} else {
			fishes[i]--
		}
	}

	return append(fishes, newFishes...)
}
