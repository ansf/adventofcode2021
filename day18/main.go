package main

import (
	"bufio"
	"os"
)

// A Node of a tree.
// If v is not -1 this node is a leaf-node, otherwise it is an inner node.
// Inner nodes have left and right child nodes.
// Leaf nodes have previous and next leaf nodes.
type Node struct {
	left, right    *Node
	next, previous *Node

	v int
}

func (n Node) Clone() Node {
	if n.v != -1 {
		return Node{v: n.v}
	}

	clone := Node{v: -1}
	left := n.left.Clone()
	right := n.right.Clone()
	clone.left = &left
	clone.right = &right

	clone.Fix()
	return clone
}

func (n Node) Magnitude() int {
	if n.v != -1 {
		return n.v
	}

	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

func (n Node) CanExplode() bool {
	return n.left != nil && n.left.v != -1 && n.right != nil && n.right.v != -1
}

// Fix restores the previous/next pointers after the tree was modified.
func (root *Node) Fix() {
	var last *Node
	VisitLeftFirst(root, 0, func(depth int, n *Node) bool {
		if n.v != -1 {
			n.previous = last
			last = n
		}
		return true
	})

	last = nil
	VisitRightFirst(root, 0, func(depth int, n *Node) bool {
		if n.v != -1 {
			n.next = last
			last = n
		}
		return true
	})

}

// Reduce reduces a single step.
// Recude returns true when the tree was modified
func (n *Node) Reduce() bool {
	cont := VisitLeftFirst(n, 0, func(depth int, node *Node) bool {
		if depth == 4 && node.CanExplode() {
			node.Explode()
			return false
		}
		return true
	})

	if !cont {
		return true
	}

	cont = VisitLeftFirst(n, 0, func(depth int, node *Node) bool {
		if node.v >= 10 {
			node.Split()
			return false
		}
		return true
	})

	return !cont
}

func (n *Node) Explode() {
	previous := n.left.previous
	if previous != nil {
		previous.v += n.left.v
	}

	next := n.right.next
	if next != nil {
		next.v += n.right.v
	}

	n.v = 0
	n.left = nil
	n.right = nil
}

func (n *Node) Split() {
	left := Node{v: n.v / 2}
	right := Node{v: (n.v + 1) / 2}
	n.left = &left
	n.right = &right
	n.v = -1
}

// Add adds two trees and recudes the result
func (n *Node) Add(m *Node) *Node {
	result := Node{
		left:  n,
		right: m,
		v:     -1,
	}

	for result.Reduce() {
		result.Fix()
	}

	return &result
}

func (n Node) Debug() {
	if n.v != -1 {
		print(n.v)
		return
	}

	print("[")
	n.left.Debug()
	print(",")
	n.right.Debug()
	print("]")
}

func (n Node) DebugNext() {
	if n.v != -1 {
		for {
			print(n.v, ",")
			if n.next == nil {
				return
			}
			n = *n.next
		}
	}

	n.left.DebugNext()
}

func (n Node) DebugPrevious() {
	if n.v != -1 {
		for {
			print(n.v, ",")
			if n.previous == nil {
				return
			}
			n = *n.previous
		}
	}

	n.right.DebugPrevious()
}

func main() {
	trees := readInput()

	max := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees); j++ {
			if i == j {
				continue
			}

			a := trees[i].Clone()
			b := trees[j].Clone()
			sum := a.Add(&b)
			mag := sum.Magnitude()
			if mag > max {
				max = mag
			}
		}
	}

	println("max magnitude: ", max)

	sum := &trees[0]
	for i := 1; i < len(trees); i++ {
		sum = sum.Add(&trees[i])
	}

	println("magnitude of sum: ", sum.Magnitude())
}

func readInput() []Node {
	result := make([]Node, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result = append(result, buildTree(scanner.Text()))
	}

	return result
}

func buildTree(input string) Node {
	reader := TreeReader{0, input}
	tree := readNode(&reader)
	tree.Fix()
	return tree
}

func VisitLeftFirst(n *Node, depth int, visitor func(depth int, node *Node) bool) bool {
	if n.v != -1 {
		return visitor(depth, n)
	}

	if !VisitLeftFirst(n.left, depth+1, visitor) {
		return false
	}
	if !visitor(depth, n) {
		return false
	}
	return VisitLeftFirst(n.right, depth+1, visitor)
}

func VisitRightFirst(n *Node, depth int, visitor func(depth int, node *Node) bool) bool {
	if n.v != -1 {
		return visitor(depth, n)
	}

	if !VisitRightFirst(n.right, depth+1, visitor) {
		return false
	}
	if !visitor(depth, n) {
		return false
	}
	return VisitRightFirst(n.left, depth+1, visitor)
}

func readNode(tr *TreeReader) Node {
	n := Node{v: -1}
	if tr.PeekNumber() {
		n.v = tr.ReadNumber()
		return n
	}

	if tr.Peek('[') {
		tr.Discard("[")
		left := readNode(tr)
		n.left = &left
		tr.Discard(",")
		right := readNode(tr)
		n.right = &right
		tr.Discard("]")
		return n
	}

	panic("parser error")
}

type TreeReader struct {
	cursor int
	input  string
}

func (tr *TreeReader) ReadNumber() int {
	if !tr.PeekNumber() {
		panic("no number")
	}

	result := 0
	for {
		if !tr.PeekNumber() {
			return result
		}

		n := int(tr.input[tr.cursor] - '0')
		tr.cursor++

		result *= 10
		result += n
	}
}

func (tr TreeReader) Peek(b byte) bool {
	return tr.input[tr.cursor] == b
}

func (tr TreeReader) PeekNumber() bool {
	b := tr.input[tr.cursor]
	return b >= '0' && b <= '9'
}

func (tr *TreeReader) Discard(s string) {
	discard := tr.input[tr.cursor : tr.cursor+len(s)]
	if s != discard {
		panic("unexpected input")
	}
	tr.cursor += len(s)
}
