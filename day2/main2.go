package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := ""
	v := 0

	depth := 0
	pos := 0
	aim := 0

	for {
		_, err := fmt.Fscanf(os.Stdin, "%s %d", &cmd, &v)
		if err != nil {
			break
		}

		if cmd == "forward" {
			pos += v
			depth += aim * v
		} else if cmd == "down" {
			aim += v
		} else if cmd == "up" {
			aim -= v
		} else {
			panic(fmt.Sprintf("unknown command %s", cmd))
		}
	}

	println(depth * pos)
}
