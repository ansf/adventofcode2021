package main

import (
	"fmt"
	"math"
	"os"
)

// just keep track of pair counts
type Template map[string]int

func ApplyInsertions(template Template, insertions map[string]string) Template {
	next := make(map[string]int)

	for pair, count := range template {
		insertion, ok := insertions[pair]
		if !ok {
			next[pair] += count
			continue
		}
		next[string(pair[0])+insertion] += count
		next[insertion+string(pair[1])] += count
	}

	return next
}

func main() {
	template, insertions := readInput()

	for step := 0; step < 40; step++ {
		template = ApplyInsertions(template, insertions)
	}

	counts := make(map[byte]int)
	for pair, count := range template {
		counts[pair[0]] += count
	}

	min := math.MaxInt
	max := 0
	for _, count := range counts {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}
	println(max - min)
}

func readInput() (Template, map[string]string) {
	var startString string
	_, err := fmt.Fscanf(os.Stdin, "%s", &startString)
	if err != nil {
		panic("could not read input")
	}

	start := make(map[string]int, 0)
	for i := 1; i < len(startString); i++ {
		start[string(startString[i-1:i+1])]++
	}
	start[string(startString[len(startString)-1])+"."]++

	fmt.Fscanln(os.Stdin)

	insertions := make(map[string]string, 0)
	for {
		var pair string
		var insertion string
		_, err = fmt.Fscanf(os.Stdin, "%s -> %s", &pair, &insertion)
		if err != nil {
			break
		}
		insertions[pair] = insertion
	}

	return start, insertions
}
