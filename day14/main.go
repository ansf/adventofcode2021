package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func ApplyInsertions(template string, insertions map[string]string) string {
	var next strings.Builder

	for i := 1; i < len(template); i++ {
		next.WriteByte(template[i-1])

		pair := string(template[i-1 : i+1])
		insertion, ok := insertions[pair]
		if ok {
			next.WriteString(insertion)
		}
	}

	next.WriteByte(template[len(template)-1])

	return next.String()
}

func main() {
	template, insertions := readInput()

	for step := 0; step < 10; step++ {
		template = ApplyInsertions(template, insertions)
	}

	counts := make(map[rune]int)
	for _, c := range template {
		counts[c]++
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

func readInput() (string, map[string]string) {
	var start string
	_, err := fmt.Fscanf(os.Stdin, "%s", &start)
	if err != nil {
		panic("could not read input")
	}

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
