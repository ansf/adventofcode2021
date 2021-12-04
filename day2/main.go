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

	for {
		_, err := fmt.Fscanf(os.Stdin, "%s %d", &cmd, &v)
		if err != nil {
			break
		}

		if cmd == "forward" {
			pos += v
		} else if cmd == "down" {
			depth += v
		} else if cmd == "up" {
			depth -= v
		} else {
			panic(fmt.Sprintf("unknown command %s", cmd))
		}
	}

	println(depth * pos)
}
