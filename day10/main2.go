package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	scores := make([]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stack := Stack{}
		input := scanner.Text()

		corrupt := false
		for _, r := range input {
			if IsOpen(r) {
				stack.Push(r)
			} else {
				expected := ClosingPair(stack.Pop())
				if expected != r {
					corrupt = true
					break
				}
			}
		}

		if !corrupt {
			scores = append(scores, ScoreString(stack.Remainder()))
		}
	}

	sort.Ints(scores)
	println(scores[len(scores)/2])
}

type Stack struct {
	data []rune
}

func (s *Stack) Push(r rune) {
	s.data = append(s.data, r)
}

func (s *Stack) Pop() rune {
	if len(s.data) == 0 {
		panic("empty stack")
	}
	r := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return r
}

func (s *Stack) Remainder() string {
	var result strings.Builder
	for {
		if len(s.data) == 0 {
			break
		}
		_, err := result.WriteRune(ClosingPair(s.Pop()))
		if err != nil {
			panic("failed to write rune")
		}
	}
	return result.String()
}

func IsOpen(r rune) bool {
	return r == '(' || r == '{' || r == '[' || r == '<'
}

func ClosingPair(r rune) rune {
	if r == '(' {
		return ')'
	} else if r == '[' {
		return ']'
	} else if r == '{' {
		return '}'
	} else if r == '<' {
		return '>'
	}
	panic("no bracket")
}

func Score(r rune) int {
	if r == ')' {
		return 3
	} else if r == ']' {
		return 57
	} else if r == '}' {
		return 1197
	} else if r == '>' {
		return 25137
	}
	panic("no bracket")
}

func ScoreClosing(r rune) int {
	if r == ')' {
		return 1
	} else if r == ']' {
		return 2
	} else if r == '}' {
		return 3
	} else if r == '>' {
		return 4
	}
	panic("no bracket")
}

func ScoreString(s string) int {
	score := 0

	for _, r := range s {
		score *= 5
		score += ScoreClosing(r)
	}

	return score
}
