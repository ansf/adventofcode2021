package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type BitStream struct {
	cursor int

	// for now simple but inefficient representation (just a binary string "0101...")
	data string
}

func NewBitStream(binary string) BitStream {
	return BitStream{0, binary}
}

func (bs BitStream) EOF() bool {
	return bs.cursor == len(bs.data)
}

func (bs *BitStream) Read(n int) uint {
	c := bs.cursor
	bs.cursor += n

	substr := string(bs.data[c : c+n])

	var result uint
	fmt.Sscanf(substr, "%b", &result)
	return result
}

func (bs *BitStream) Stream(n int) string {
	c := bs.cursor
	bs.cursor += n
	return string(bs.data[c : c+n])
}

const (
	packetTypeSum     = 0
	packetTypeProduct = 1
	packetTypeMin     = 2
	packetTypeMax     = 3
	packetTypeLiteral = 4
	packetTypeGT      = 5
	packetTypeLT      = 6
	packetTypeEQ      = 7

	lengthTypeBits    = 0
	lengthTypePackets = 1
)

type Packet struct {
	Type    int
	Version int

	Value int

	Packets []Packet
}

type Lexer struct {
	bs *BitStream
}

func (l Lexer) ReadVersion() int {
	return int(l.bs.Read(3))
}

func (l Lexer) ReadType() int {
	return int(l.bs.Read(3))
}

func (l Lexer) ReadLiteral() int {
	var result uint
	for {
		i := l.bs.Read(1)
		result <<= 4
		result |= l.bs.Read(4)
		if i == 0 {
			break
		}
	}
	return int(result)
}

func (l Lexer) ReadOperatorLength() (int, int) {
	t := int(l.bs.Read(1))
	switch t {
	case lengthTypeBits:
		return t, int(l.bs.Read(15))
	case lengthTypePackets:
		return t, int(l.bs.Read(11))
	default:
		panic("failed to parse operator length")
	}
}

type Parser struct {
	l Lexer
}

func NewParser(bs *BitStream) Parser {
	return Parser{Lexer{bs}}
}

func (p Parser) Parse() Packet {
	v := p.l.ReadVersion()
	t := p.l.ReadType()

	switch t {
	case packetTypeLiteral:
		value := p.l.ReadLiteral()
		return Packet{Type: t, Version: v, Value: value}

	default:
		lengthType, length := p.l.ReadOperatorLength()
		switch lengthType {
		case lengthTypeBits:
			bs := NewBitStream(p.l.bs.Stream(length))
			parser := NewParser(&bs)
			result := make([]Packet, 0)
			for !bs.EOF() {
				result = append(result, parser.Parse())
			}
			return Packet{Type: t, Version: v, Packets: result}

		case lengthTypePackets:
			result := make([]Packet, 0)
			for i := 0; i < length; i++ {
				result = append(result, p.Parse())
			}
			return Packet{Type: t, Version: v, Packets: result}

		default:
			panic("unknown lengthType")
		}
	}
}

func SumVersion(packet Packet) int {
	sum := packet.Version
	for _, p := range packet.Packets {
		sum += SumVersion(p)
	}
	return sum
}

func Evaluate(packet Packet) int {
	switch packet.Type {
	case packetTypeSum:
		sum := 0
		for _, p := range packet.Packets {
			sum += Evaluate(p)
		}
		return sum
	case packetTypeProduct:
		sum := 1
		for _, p := range packet.Packets {
			sum *= Evaluate(p)
		}
		return sum
	case packetTypeMin:
		min := math.MaxInt
		for _, p := range packet.Packets {
			v := Evaluate(p)
			if v < min {
				min = v
			}
		}
		return min
	case packetTypeMax:
		max := math.MinInt
		for _, p := range packet.Packets {
			v := Evaluate(p)
			if v > max {
				max = v
			}
		}
		return max
	case packetTypeLiteral:
		return packet.Value
	case packetTypeGT:
		v1 := Evaluate(packet.Packets[0])
		v2 := Evaluate(packet.Packets[1])
		if v1 > v2 {
			return 1
		}
		return 0
	case packetTypeLT:
		v1 := Evaluate(packet.Packets[0])
		v2 := Evaluate(packet.Packets[1])
		if v1 < v2 {
			return 1
		}
		return 0
	case packetTypeEQ:
		v1 := Evaluate(packet.Packets[0])
		v2 := Evaluate(packet.Packets[1])
		if v1 == v2 {
			return 1
		}
		return 0
	default:
		panic("unknown packet type")
	}
}

func main() {
	var input string
	_, err := fmt.Fscanf(os.Stdin, "%s", &input)
	if err != nil {
		log.Fatal(err.Error())
	}

	bs := NewBitStream(hexToBinary(input))

	parser := Parser{Lexer{&bs}}
	p := parser.Parse()

	println(SumVersion(p))

	println(Evaluate(p))
}

func hexToBinary(hex string) string {
	reader := bytes.NewBufferString(hex)

	var data strings.Builder
	for i := 0; i < len(hex)/2; i++ {
		var b uint8
		fmt.Fscanf(reader, "%2X", &b)
		data.WriteString(fmt.Sprintf("%08b", b))
	}

	return data.String()
}
