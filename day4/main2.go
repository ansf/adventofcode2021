package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	drawing := readDrawing(scanner)
	boards := readBoards(scanner)

	remaining := len(boards)

	for _, d := range drawing {
		for _, b := range boards {
			if b.Won() {
				continue
			}

			b.Draw(d)
			if b.Won() {
				remaining--
				if remaining == 0 {
					println(d * b.Score())
					return
				}
			}
		}
	}

}

type Board struct {
	data [][]int
	hit  [][]bool
}

func NewBoard() *Board {
	board := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	hit := [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}
	return &Board{board, hit}
}

func (b *Board) Draw(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.data[i][j] == n {
				b.hit[i][j] = true
			}
		}
	}
}

func (b *Board) Won() bool {
	for i := 0; i < 5; i++ {
		if b.hit[i][0] && b.hit[i][1] && b.hit[i][2] && b.hit[i][3] && b.hit[i][4] {
			return true
		}
		if b.hit[0][i] && b.hit[1][i] && b.hit[2][i] && b.hit[3][i] && b.hit[4][i] {
			return true
		}
	}
	return false
}

func (b *Board) Score() int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.hit[i][j] {
				score += b.data[i][j]
			}
		}
	}
	return score
}

func (b *Board) String() string {
	s := fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n", b.data[0], b.data[1], b.data[2], b.data[3], b.data[4])
	s += fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n", b.hit[0], b.hit[1], b.hit[2], b.hit[3], b.hit[4])
	return s
}

func readDrawing(scanner *bufio.Scanner) []int {
	if !scanner.Scan() {
		panic("could not read drawing")
	}

	s := scanner.Text()
	ss := strings.Split(s, ",")

	result := make([]int, len(ss))
	for i := 0; i < len(result); i++ {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			panic("could not parse numbers")
		}

		result[i] = n
	}

	return result
}

func readBoards(scanner *bufio.Scanner) []*Board {
	boards := make([]*Board, 0)
	for scanner.Scan() {
		boards = append(boards, readBoard(scanner))
	}
	return boards
}

func readBoard(scanner *bufio.Scanner) *Board {
	board := NewBoard()

	for i := 0; i < 5; i++ {
		if !scanner.Scan() {
			panic("could not read boards")
		}

		n, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d %d",
			&board.data[i][0],
			&board.data[i][1],
			&board.data[i][2],
			&board.data[i][3],
			&board.data[i][4],
		)
		if n != 5 || err != nil {
			panic("could not parse boards")
		}
	}

	return board
}
