package lib

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ReadListOfNumbers(r io.Reader) ([]int, error) {
	input := ""
	_, err := fmt.Fscanf(r, "%s", &input)
	if err != nil {
		return nil, err
	}

	positions := make([]int, 0)
	for _, v := range strings.Split(input, ",") {
		p, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		positions = append(positions, p)
	}

	return positions, nil
}
