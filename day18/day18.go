package day18

import (
	"AdventCode2021/util"
	"fmt"
	"math"
)

const (
	TypeValue = 1
	TypePair  = 2
)

func Resolve(input []string) (string, int) {
	sum := parsePair(&stream{input: input[0]})
	for i := 1; i < len(input); i++ {
		sum = &pair{x: newPairValue(sum), y: parsePairValue(input[i])}
		reduce(sum)
	}

	reduce(sum)
	return sum.string(), sum.magnitude()
}

func LargestMagnitude(input []string) int {
	largest := math.MinInt
	for _, i := range input {
		for _, j := range input {
			if i == j {
				continue
			}

			_, mag := Resolve([]string{i, j})
			if mag > largest {
				largest = mag
			}
		}
	}

	return largest
}

func reduce(sum *pair) {
	for {
		data := &explodedData{}
		sum.explode(0, data)
		if data.exploded != nil {
			continue
		}

		if sum.split() == false {
			break
		}
	}
}

type stream struct {
	input string
	pos   int
}

func (p *stream) pop() string {
	if len(p.input) == 0 {
		return ""
	}
	char := p.input[:1]
	p.input = p.input[1:]
	return char
}

func (p *stream) peek() string {
	return p.input[0:1]
}

type value struct {
	t int
	p *pair
	v int
}

func (v value) string() string {
	if v.t == TypeValue {
		return fmt.Sprintf("%d", v.v)
	}

	return v.p.string()
}

func (v value) magnitude() int {
	if v.t == TypeValue {
		return v.v
	}

	return v.p.magnitude()
}

func newPairValue(p *pair) value {
	return value{t: TypePair, p: p}
}

func parsePairValue(l string) value {
	return newPairValue(parsePair(&stream{input: l}))
}

type pair struct {
	x value
	y value
}

func (p *pair) isValuePair() bool {
	return p.x.t == TypeValue && p.y.t == TypeValue
}

func (p *pair) magnitude() int {
	return p.x.magnitude()*3 + p.y.magnitude()*2
}

func (p *pair) string() string {
	return fmt.Sprintf("[%s,%s]", p.x.string(), p.y.string())
}

func (p *pair) split() bool {
	//TOD: Move code in to value
	if p.x.t == TypeValue && p.x.v >= 10 {
		p.x.t = TypePair
		p.x.p = &pair{
			x: value{t: TypeValue, v: int(math.Floor(float64(p.x.v) / 2.0))},
			y: value{t: TypeValue, v: int(math.Ceil(float64(p.x.v) / 2.0))},
		}
		return true
	} else if p.x.t == TypePair {
		split := p.x.p.split()
		if split {
			return true
		}
	}

	if p.y.t == TypeValue && p.y.v >= 10 {
		p.y.t = TypePair
		p.y.p = &pair{
			x: value{t: TypeValue, v: int(math.Floor(float64(p.y.v) / 2.0))},
			y: value{t: TypeValue, v: int(math.Ceil(float64(p.y.v) / 2.0))},
		}
		return true
	} else if p.y.t == TypePair {
		split := p.y.p.split()
		if split {
			return true
		}
	}

	return false
}

type explodedData struct {
	prev     *value
	next     *value
	exploded *pair
}

func (p *pair) explode(depth int, data *explodedData) {

	if p.isValuePair() && depth == 4 && data.next == nil && data.exploded == nil {
		if data.prev != nil {
			data.prev.v += p.x.v
		}
		data.next = &p.y
		data.exploded = p
	}

	//TOD: Move code in to value
	if p.x.t == TypeValue {
		if data.exploded != p && data.next != nil {
			//fmt.Printf("Adding %d on to %d\n", data.next.v, p.x.v)
			p.x.v += data.next.v
			data.next = nil
		}
		data.prev = &p.x
	} else {
		p.x.p.explode(depth+1, data)
	}

	if p.y.t == TypeValue {
		if data.exploded != p && data.next != nil {
			p.y.v += data.next.v
			data.next = nil
		}
		data.prev = &p.y
	} else {
		p.y.p.explode(depth+1, data)
	}

	if data.exploded != nil && p.x.p == data.exploded {
		p.x.t = TypeValue
		p.x.v = 0
	} else if data.exploded != nil && p.y.p == data.exploded {
		p.y.t = TypeValue
		p.y.v = 0
	}
}

func parsePair(s *stream) *pair {
	p := &pair{}
	c := s.pop()

	//TOD: create pop assert check method
	if c != "[" {
		panic(fmt.Sprintf("invalid character, expected [ was %s", c))
	}

	p.x = newValue(s)

	c = s.pop()
	if c != "," {
		panic(fmt.Sprintf("invalid character, expected , was %s", c))
	}

	p.y = newValue(s)

	c = s.pop()
	if c != "]" {
		panic(fmt.Sprintf("invalid character, expected ] was %s", c))
	}

	return p
}

func newValue(s *stream) value {
	c := s.peek()
	if c == "[" {
		return value{t: TypePair, p: parsePair(s)}
	}

	return value{t: TypeValue, v: util.TryParseInt(s.pop())}
}
