package day16

import (
	"AdventCode2021/util"
	"strings"
)

const (
	LengthVersion         = 3
	LengthType            = 3
	LengthLiteralChunk    = 5
	LengthOpLengthType    = 1
	LengthOpPacketsSize   = 15
	LengthOpPacketsNumber = 11

	TypeSum         = 0
	TypeProduct     = 1
	TypeMin         = 2
	TypeMax         = 3
	TypeGreaterThan = 5
	TypeLessThan    = 6
	TypeEqual       = 7
	TypeLiteral     = 4
)

type stream struct {
	binary string
	pos    int
}

func (p *stream) next(n int) string {
	bits := p.binary[:n]
	p.binary = p.binary[n:]
	return bits
}

func (p *stream) nextInt(n int) int {
	return util.TryParseBinary(p.next(n))
}

func (p *stream) remaining() int {
	return len(p.binary)
}

func (p *stream) hasBits() bool {
	return p.remaining() > 0 && strings.Contains(p.binary, "1")
}

func ParseHex(input string) (int, int64) {
	s := &stream{binary: getBinaryString(input)}
	sum, vs := parsePacket(s, 0)
	return sum, vs[0]
}

func parsePacket(s *stream, limit int) (int, []int64) {
	sum, packets, values := 0, 0, make([]int64, 0)
	for s.hasBits() {
		v, t := s.nextInt(LengthVersion), s.nextInt(LengthType)
		sum += v
		if t == TypeLiteral {
			values = append(values, parseLiteral(s))
		} else {
			var subValues []int64
			var subSum int
			if s.nextInt(LengthOpLengthType) == 0 {
				subSum, subValues = parsePacket(&stream{binary: s.next(s.nextInt(LengthOpPacketsSize))}, 0)
			} else {
				subSum, subValues = parsePacket(s, s.nextInt(LengthOpPacketsNumber))
			}

			sum += subSum
			values = append(values, applyOperator(t, subValues))
		}

		packets++
		if limit > 0 && packets == limit {
			break
		}
	}

	return sum, values
}

func applyOperator(t int, subValues []int64) int64 {
	operatedValue := int64(0)
	switch t {
	case TypeSum:
		operatedValue = util.Sum64(subValues)
	case TypeProduct:
		operatedValue = util.Product64(subValues)
	case TypeMin:
		operatedValue = util.Min64(subValues)
	case TypeMax:
		operatedValue = util.Max64(subValues)
	case TypeGreaterThan:
		if subValues[0] > subValues[1] {
			operatedValue = 1
		} else {
			operatedValue = 0
		}
	case TypeLessThan:
		if subValues[0] < subValues[1] {
			operatedValue = 1
		} else {
			operatedValue = 0
		}
	case TypeEqual:
		if subValues[0] == subValues[1] {
			operatedValue = 1
		} else {
			operatedValue = 0
		}
	}
	return operatedValue
}

func parseLiteral(s *stream) int64 {
	var bits string
	for {
		b := s.next(LengthLiteralChunk)
		bits += b[1:LengthLiteralChunk]
		if string(b[0]) == "0" {
			break
		}
	}
	return util.TryParseBinary64(bits)
}

func getBinaryString(hex string) string {
	mapping := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	var sb strings.Builder
	for _, h := range hex {
		sb.WriteString(mapping[string(h)])
	}

	return sb.String()
}
