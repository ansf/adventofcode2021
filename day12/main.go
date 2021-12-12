package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Path []string
type Caves map[string][]string

func FindNext(caves Caves, path Path) []string {
	result := Path{}
	for _, to := range caves[path[len(path)-1]] {
		if IsSmallCave(to) && Contains(path, to) {
			continue
		}
		result = append(result, to)
	}
	return result
}

func FindAll(caves Caves, start string) []Path {
	validPaths := make([]Path, 0)

	paths := []Path{{start}}
	nextPaths := make([]Path, 0)

	for {
		for _, path := range paths {
			nextCaves := FindNext(caves, path)
			for _, nextCave := range nextCaves {
				nextPath := make(Path, len(path))
				copy(nextPath, path)
				nextPath = append(nextPath, nextCave)

				if nextCave == "end" {
					validPaths = append(validPaths, nextPath)
				}

				nextPaths = append(nextPaths, nextPath)
			}
		}

		if len(nextPaths) == 0 {
			break
		}

		paths = nextPaths
		nextPaths = make([]Path, 0)
	}

	return validPaths
}

func IsSmallCave(name string) bool {
	return name[0] > 'Z'
}

func Contains(path Path, name string) bool {
	for _, p := range path {
		if name == p {
			return true
		}
	}
	return false
}

func main() {
	caves := readInput()
	fmt.Printf("%v\n", len(FindAll(caves, "start")))
}

func readInput() Caves {
	caves := make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fromTo := strings.Split(input, "-")
		caves[fromTo[0]] = append(caves[fromTo[0]], fromTo[1])
		caves[fromTo[1]] = append(caves[fromTo[1]], fromTo[0])
	}

	return caves
}
