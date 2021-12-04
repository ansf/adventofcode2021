package main

import (
	"fmt"
	"os"
)

func main() {
	a := read()

	current := a[0] + a[1] + a[2]
	last := current

	count := 0

	for i := 3; i < len(a); i++ {
		current -= a[i-3]
		current += a[i]

		if current > last {
			count++
		}

		last = current
	}

	fmt.Println(count)
}

func read() []int {
	a := make([]int, 0)

	for i := 0; ; i++ {
		var v int
		_, err := fmt.Fscanf(os.Stdin, "%d", &v)
		if err != nil {
			break
		}
		a = append(a, v)
	}

	return a
}
