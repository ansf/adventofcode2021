package main

import (
	"fmt"
	"os"
)

func main() {
	v := ""

	var countOnes []int = nil

	lines := 0
	for ; ; lines++ {
		_, err := fmt.Fscanf(os.Stdin, "%s", &v)
		if err != nil {
			break
		}

		if countOnes == nil {
			countOnes = make([]int, len(v))
		}

		for i := 0; i < len(v); i++ {
			if v[i] == '1' {
				countOnes[i]++
			}
		}
	}

	gamma := 0
	epsilon := 0
	for i := 0; i < len(countOnes); i++ {
		v := 0
		if countOnes[i] > lines/2 {
			v = 1
		}
		gamma = (gamma << 1) | v
		epsilon = (epsilon << 1) | (^v & 1)
	}

	println(gamma * epsilon)
}
