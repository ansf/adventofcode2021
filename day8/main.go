package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	outputs := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		segments := strings.Split(line, " | ")
		outputs = append(outputs, strings.Split(segments[1], " "))
	}

	count := 0
	for _, output := range outputs {
		for _, digit := range output {
			length := len(digit)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				count++
			}
		}
	}

	println(count)
}
