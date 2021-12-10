package main

import (
	"bufio"
	"os"
)

func main() {
	score := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stack := Stack{}
		input := scanner.Text()

		for _, r := range input {
			if IsOpen(r) {
				stack.Push(r)
			} else {
				expected := ClosingPair(stack.Pop())
				if expected != r {
					score += Score(r)
					break
				}
			}
		}
	}

	println(score)
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
