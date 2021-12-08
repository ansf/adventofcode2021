package main

import (
	"bufio"
	"os"
	"strings"
)

var segments = [][]bool{
	{true, true, true, false, true, true, true},
	{false, false, true, false, false, true, false},
	{true, false, true, true, true, false, true},
	{true, false, true, true, false, true, true},
	{false, true, true, true, false, true, false},
	{true, true, false, true, false, true, true},
	{true, true, false, true, true, true, true},
	{true, false, true, false, false, true, false},
	{true, true, true, true, true, true, true},
	{true, true, true, true, false, true, true},
}

func PossibleDigits(l int) []int {
	switch l {
	case 2:
		return []int{1}
	case 3:
		return []int{7}
	case 4:
		return []int{4}
	case 7:
		return []int{8}
	case 6:
		return []int{0, 6, 9}
	case 5:
		return []int{2, 3, 5}
	default:
		panic("unexpected length")
	}
}

// abc... => 0,1,2...
func Atoi(s string) []int {
	result := make([]int, 0)
	for i := 0; i < len(s); i++ {
		result = append(result, int(s[i]-'a'))
	}
	return result
}

func Contains(a []int, n int) bool {
	for _, v := range a {
		if n == v {
			return true
		}
	}
	return false
}

type Wire struct {
	on []bool
}

func NewWire() Wire {
	return Wire{[]bool{true, true, true, true, true, true, true}}
}

func (w *Wire) IsAnyOn(in []int) bool {
	for _, i := range in {
		if w.on[i] {
			return true
		}
	}
	return false
}

// SwitchOthersOff returns true if something was actually switched
func (w *Wire) SwitchOthersOff(in []int) bool {
	result := false
	for i := 0; i < len(w.on); i++ {
		if !Contains(in, i) {
			if w.on[i] {
				result = true
			}
			w.on[i] = false
		}
	}

	return result
}

// SwitchOff returns true if something was actually switched
func (w *Wire) SwitchOff(in []int) bool {
	result := false

	for i := 0; i < len(w.on); i++ {
		if Contains(in, i) {
			if w.on[i] {
				result = true
			}
			w.on[i] = false
		}
	}

	return result
}

// Decuder starts with every wire possibly going to any segment
type Deducer struct {
	wiring []Wire
}

func NewDeducer() Deducer {
	return Deducer{
		[]Wire{
			NewWire(),
			NewWire(),
			NewWire(),
			NewWire(),
			NewWire(),
			NewWire(),
			NewWire(),
		},
	}
}

// Learn will reduce the wiring possibilities for the given input
func (d *Deducer) Learn(s string) bool {

	// if a segment is on for every possible digit of the given input,
	// we know that we can exclude all other wirings for this segment.
	// if a segment is on for none of the possible digits of the given input,
	// we know that we can exclude all given wirings for this segment.

	possibleDigits := d.Decode(s)
	inputWires := Atoi(s)

	allOn := []bool{true, true, true, true, true, true, true}
	someOn := []bool{false, false, false, false, false, false, false}
	for _, digit := range possibleDigits {
		segmentsOn := segments[digit]
		for i := 0; i < len(segmentsOn); i++ {
			allOn[i] = allOn[i] && segmentsOn[i]
			someOn[i] = someOn[i] || segmentsOn[i]
		}
	}

	learned := false
	for i := 0; i < len(allOn); i++ {
		if allOn[i] {
			learned = d.wiring[i].SwitchOthersOff(inputWires) || learned
		} else if !someOn[i] {
			learned = d.wiring[i].SwitchOff(inputWires) || learned
		}
	}

	return learned
}

// Decode given input into all possible digits, using the
// current wiring possibilities of the deducer
func (d *Deducer) Decode(s string) []int {
	possibleDigits := PossibleDigits(len(s))
	inputWires := Atoi(s)

	// rule out digits which cannot be wired
	validDigits := make([]int, 0)
	for _, digit := range possibleDigits {
		segmentsOn := segments[digit]
		valid := true
		for i := 0; i < 7; i++ {
			if segmentsOn[i] && !d.wiring[i].IsAnyOn(inputWires) {
				valid = false
				break
			}
		}
		if valid {
			validDigits = append(validDigits, digit)
		}
	}

	if len(validDigits) == 0 {
		panic("no valid digits left :(")
	}

	return validDigits
}

func (d *Deducer) MustDecode(s string) int {
	decoded := d.Decode(s)
	if len(decoded) != 1 {
		panic("cannot decode")
	}
	return decoded[0]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		segments := strings.Split(line, " | ")
		inputs := strings.Split(segments[0], " ")
		outputs := strings.Split(segments[1], " ")

		deducer := NewDeducer()

		// learn from both inputs and outputs until we gain no more knowledge
		for {
			learning := false
			for _, input := range inputs {
				if deducer.Learn(input) {
					learning = true
				}
			}
			for _, input := range outputs {
				if deducer.Learn(input) {
					learning = true
				}
			}
			if !learning {
				break
			}
		}

		// decode with learned deducer
		decodedValue := 0
		for _, o := range outputs {
			decodedValue *= 10
			decodedValue += deducer.MustDecode(o)
		}

		sum += decodedValue
	}

	println(sum)
}
