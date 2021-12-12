package main

import (
	"bufio"
	"os"
	"strings"
)

type Path []string
type Caves map[string][]string

func FindNext(caves Caves, path Path, visitTwice string) []string {
	result := Path{}
	for _, to := range caves[path[len(path)-1]] {
		if IsSmallCave(to) {
			visited := Contained(path, to)
			if (visitTwice == to && visited == 2) ||
				(visitTwice != to && visited == 1) {
				continue
			}
		}
		result = append(result, to)
	}
	return result
}

func FindAll(caves Caves, start string, visitTwice string) []Path {
	validPaths := make([]Path, 0)

	paths := []Path{{start}}
	nextPaths := make([]Path, 0)

	for {
		for _, path := range paths {
			nextCaves := FindNext(caves, path, visitTwice)
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

func Contained(path Path, name string) int {
	count := 0
	for _, p := range path {
		if name == p {
			count++
		}
	}
	return count
}

func main() {
	caves := readInput()

	smallCaves := make([]string, 0)
	for cave := range caves {
		if cave != "start" && cave != "end" && IsSmallCave(cave) {
			smallCaves = append(smallCaves, cave)
		}
	}

	paths := make([]Path, 0)
	for _, smallCave := range smallCaves {
		foundPaths := FindAll(caves, "start", smallCave)
		for _, foundPath := range foundPaths {
			paths = append(paths, foundPath)
		}
	}

	uniquePaths := make(map[string]struct{}, 0)
	for _, path := range paths {
		uniquePaths[strings.Join(path, ",")] = struct{}{}
	}

	println(len(uniquePaths))
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
