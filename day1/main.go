package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	current := 0
	last := math.MaxInt

	count := 0

	for {
		_, err := fmt.Fscanf(os.Stdin, "%d", &current)
		if err != nil {
			break
		}

		if current > last {
			count++
		}

		last = current
	}

	fmt.Println(count)
}
