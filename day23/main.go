package main

import (
	"math"
	"strings"
)

const (
	roomTypeHallway = 4
)

type Burrow struct {
	cost      int
	amphipods []Amphipod
}

func (b Burrow) RoomSize() int {
	return len(b.amphipods) / 4
}

func (b Burrow) findColor(roomType int, position int) int {
	for _, a := range b.amphipods {
		if a.roomType == roomType && a.position == position {
			return a.color
		}
	}
	return -1
}

func color(color int) string {
	if color == -1 {
		return "."
	} else {
		return string(rune('A' + color))
	}
}

func (b Burrow) Key() string {
	var builder strings.Builder

	for i := 0; i < 11; i++ {
		builder.WriteString(color(b.findColor(roomTypeHallway, i)))
	}

	for i := 0; i < b.RoomSize(); i++ {
		builder.WriteString(color(b.findColor(0, i)))
		builder.WriteString(color(b.findColor(1, i)))
		builder.WriteString(color(b.findColor(2, i)))
		builder.WriteString(color(b.findColor(3, i)))
	}

	return builder.String()
}

func (b Burrow) Copy() Burrow {
	result := Burrow{b.cost, make([]Amphipod, len(b.amphipods))}
	for i := 0; i < len(b.amphipods); i++ {
		result.amphipods[i] = b.amphipods[i]
	}
	return result
}

func (b Burrow) canEnterHallway(a Amphipod) bool {
	if a.roomType == roomTypeHallway {
		// already in hallway
		return true
	}
	if a.position == 0 {
		// in upper room position
		return true
	}
	for p := a.position - 1; p >= 0; p-- {
		if b.findColor(a.roomType, p) != -1 {
			// upper room position blocked
			return false
		}
	}
	return true
}

func (b Burrow) enterRoom(a Amphipod) int {
	if a.roomType == a.color {
		return -1
	}

	roomPosition := b.targetFreeRoomPosition(a.color)
	// target room occupied
	if roomPosition == -1 {
		return -1
	}

	currentHallwayPosition := a.currentHallwayPosition()
	targetHallwayPosition := a.targetHallwayPosition()

	// hallway blocked
	if !b.isHallwayFree(a, a.currentHallwayPosition(), a.targetHallwayPosition()) {
		return -1
	}

	// upper room position blocked
	if !b.canEnterHallway(a) {
		return -1
	}

	// move into target room
	moves := (roomPosition + 1)
	// move within hallway
	moves += abs(currentHallwayPosition - targetHallwayPosition)
	if a.roomType != roomTypeHallway {
		// move into hallway
		moves += a.position + 1
	}

	return moves
}

func (b Burrow) isInFinalPosition(a Amphipod) bool {
	if a.roomType != a.color {
		return false
	}

	if a.position == b.RoomSize()-1 {
		return true
	}

	for p := a.position + 1; p < b.RoomSize(); p++ {
		// wrong color below amphoid?
		if b.findColor(a.roomType, p) != a.color {
			return false
		}
	}

	return true
}

func (b Burrow) reachableHallwayPositions(a Amphipod) []int {
	if !b.canEnterHallway(a) {
		return []int{}
	}

	if a.roomType == roomTypeHallway {
		// dont move within hallway
		return []int{}
	}

	if b.isInFinalPosition(a) {
		return []int{}
	}

	result := make([]int, 0)
	current := a.currentHallwayPosition()
	for i := 0; i < 11; i++ {
		if i == 2 || i == 4 || i == 6 || i == 8 {
			continue
		}
		if b.isHallwayFree(a, i, current) {
			result = append(result, i)
		}
	}
	return result
}

func (b Burrow) isSolution() bool {
	for _, a := range b.amphipods {
		if a.roomType != a.color {
			return false
		}
	}
	return true
}

func (b Burrow) isHallwayFree(a Amphipod, from, to int) bool {
	if from > to {
		from, to = to, from
	}
	for _, other := range b.amphipods {
		if other.id == a.id {
			continue
		}
		if other.roomType == roomTypeHallway && other.position >= from && other.position <= to {
			return false
		}
	}
	return true
}

func (b Burrow) targetFreeRoomPosition(roomType int) int {
	for p := 0; p < b.RoomSize(); p++ {
		// wrong amphipod in room, can't enter
		c := b.findColor(roomType, p)
		if c != -1 && c != roomType {
			return -1
		}
	}

	for p := b.RoomSize() - 1; p >= 0; p-- {
		// find lowest free position
		if b.findColor(roomType, p) == -1 {
			return p
		}
	}

	return -1
}

type Amphipod struct {
	id       int
	color    int
	position int
	roomType int
}

func (a Amphipod) currentHallwayPosition() int {
	if a.roomType == roomTypeHallway {
		return a.position
	}
	return 2 + a.roomType*2
}

func (a Amphipod) targetHallwayPosition() int {
	return 2 + a.color*2
}

func applyPossibleMoves(amphipodIdx int, burrow Burrow) []Burrow {
	amphipod := burrow.amphipods[amphipodIdx]
	if burrow.isInFinalPosition(amphipod) {
		return []Burrow{}
	}

	result := make([]Burrow, 0)

	moveIntoTargetRoom := burrow.enterRoom(amphipod)

	if moveIntoTargetRoom != -1 {
		b := burrow.Copy()
		b.cost += moveIntoTargetRoom * cost(amphipod.color)
		b.amphipods[amphipodIdx] = Amphipod{
			amphipod.id,
			amphipod.color,
			burrow.targetFreeRoomPosition(amphipod.color),
			amphipod.color,
		}

		result = append(result, b)
	}

	movesIntoHallway := 0
	if amphipod.roomType != roomTypeHallway {
		// move into hallway
		movesIntoHallway = amphipod.position + 1
	}

	hallwayPositions := burrow.reachableHallwayPositions(amphipod)
	for _, target := range hallwayPositions {
		current := amphipod.currentHallwayPosition()

		b := burrow.Copy()
		b.cost += (movesIntoHallway + abs(current-target)) * cost(amphipod.color)
		b.amphipods[amphipodIdx] = Amphipod{
			amphipod.id,
			amphipod.color,
			target,
			roomTypeHallway,
		}
		result = append(result, b)
	}

	return result
}

var seenBurrows = make(map[string]int, 0)
var solution = Burrow{math.MaxInt, []Amphipod{}}

func solve(burrow Burrow) {
	if burrow.isSolution() {
		if burrow.cost < solution.cost {
			solution = burrow
		}
		return
	}

	for i := 0; i < len(burrow.amphipods); i++ {
		nextBurrows := applyPossibleMoves(i, burrow)
		for _, next := range nextBurrows {
			if next.cost >= solution.cost {
				continue
			}

			cost, seen := seenBurrows[next.Key()]
			if !seen || next.cost < cost {
				seenBurrows[next.Key()] = next.cost
				solve(next)
			}
		}
	}
}

func main() {
	// input test 1: BCBDADCA
	// input test 2: BCBDDCBADBACADCA
	// input 1: DCDBBAAC
	// input 2: DCDBDCBADBACBAAC

	burrow := createBurrow("DCDBDCBADBACBAAC")
	solve(burrow)
	println(solution.cost)
}

func createBurrow(input string) Burrow {
	pods := make([]Amphipod, len(input))
	result := Burrow{0, pods}

	for i := 0; i < len(input); i++ {
		room := i % 4
		position := i / 4
		color := int(input[i] - 'A')
		result.amphipods[i] = Amphipod{i, color, position, room}
	}

	return result
}

func cost(color int) int {
	if color == 0 {
		return 1
	} else if color == 1 {
		return 10
	} else if color == 2 {
		return 100
	}
	return 1000
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
