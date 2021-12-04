package main

import (
	"fmt"
	"os"
)

func main() {
	values := make([]int, 0)

	lines := 0
	width := 0
	for ; ; lines++ {
		v := ""
		_, err := fmt.Fscanf(os.Stdin, "%s", &v)
		if err != nil {
			break
		}

		width = len(v)

		value := 0
		_, err = fmt.Sscanf(v, "%b", &value)
		if err != nil {
			panic("invalid input")
		}

		values = append(values, value)
	}

	oxy := reduce(values, width, filterToMostCommonBit)
	co2 := reduce(values, width, filterToLeastCommonBit)
	println(oxy * co2)
}

func reduce(a []int, width int, filter func([]int, int) []int) int {
	reduced := make([]int, len(a))
	copy(reduced, a)

	for i := width - 1; i >= 0; i-- {
		reduced = filter(reduced, i)

		if len(reduced) == 1 {
			return reduced[0]
		}
	}
	panic("could not reduce to single value")
}

func countOnes(a []int, idx int) int {
	countOnes := 0
	for i := 0; i < len(a); i++ {
		if (a[i]>>idx)&1 == 1 {
			countOnes++
		}
	}
	return countOnes
}

func filterBit(a []int, idx int, bit int) []int {
	result := make([]int, 0, len(a))
	for i := 0; i < len(a); i++ {
		if (a[i]>>idx)&1 == bit {
			result = append(result, a[i])
		}
	}
	return result

}

func filterToMostCommonBit(a []int, idx int) []int {
	countOnes := countOnes(a, idx)

	mostCommon := 0
	if countOnes*2 >= len(a) {
		mostCommon = 1
	}

	return filterBit(a, idx, mostCommon)
}

func filterToLeastCommonBit(a []int, idx int) []int {
	countOnes := countOnes(a, idx)

	leastCommon := 1
	if countOnes*2 >= len(a) {
		leastCommon = 0
	}

	return filterBit(a, idx, leastCommon)
}
