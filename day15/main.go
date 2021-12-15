package main

import (
	"bufio"
	"container/heap"
	"math"
	"os"
)

type Point struct {
	x, y int
}

func (p Point) Adjacents() []Point {
	return []Point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y - 1},
		{p.x, p.y + 1},
	}
}

type Node struct {
	priority int

	point Point
	prev  *Node

	index int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	// get last item
	old := *pq
	n := len(old)
	item := old[n-1]

	// shorten slice
	old[n-1] = nil
	*pq = old[0 : n-1]

	return item
}

func (pq PriorityQueue) find(p Point) *Node {
	for _, node := range pq {
		if node.point == p {
			return node
		}
	}
	return nil
}

func (pq PriorityQueue) neighbors(p Point) []*Node {
	result := make([]*Node, 0)
	for _, adjacent := range p.Adjacents() {
		node := pq.find(adjacent)
		if node != nil {
			result = append(result, node)
		}
	}

	return result
}

func (pq *PriorityQueue) update(node *Node, priority int, prev *Node) {
	node.priority = priority
	node.prev = prev
	heap.Fix(pq, node.index)
}

func CheapestPath(graph [][]int) int {
	pq := make(PriorityQueue, 0)
	i := 0
	for y, g := range graph {
		for x := range g {
			pq = append(pq, &Node{
				priority: math.MaxInt,
				point:    Point{x, y},
				index:    i,
			})
			i++
		}
	}

	// start node
	pq[0].priority = 0
	heap.Init(&pq)

	w := len(graph[0])
	h := len(graph)

	for len(pq) > 0 {
		node := heap.Pop(&pq).(*Node)

		if node.point.x == w-1 && node.point.y == h-1 {
			return node.priority
		}

		neighbors := pq.neighbors(node.point)
		for _, neighbor := range neighbors {
			d := node.priority + graph[neighbor.point.y][neighbor.point.x]
			if d < neighbor.priority {
				pq.update(neighbor, d, node)
			}
		}
	}

	visited := make([][]bool, 0)
	for range graph {
		v := make([]bool, len(graph[0]))
		visited = append(visited, v)
	}

	return 0
}

func main() {
	input := readInput()
	println(CheapestPath(input))
}

func readInput() [][]int {
	result := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		line := make([]int, len(input))
		for i := 0; i < len(line); i++ {
			line[i] = int(input[i] - '0')
		}

		result = append(result, line)
	}

	return result
}
