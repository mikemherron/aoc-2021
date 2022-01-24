package alu

import (
	"AdventCode2021/util"
	"fmt"
	"strings"
)

func NewALU() *ALU {
	// Create Z as a property on struct to avoid map look up
	// on each iteration from monad program
	return &ALU{}
}

type ALU struct {
	X, Y, Z, W int
}

type Program struct {
	s      []*statement
	labels map[string][]int
}

const Label = ":"

type instruction func(a *int, b *int, i *input, c *int)

type statement struct {
	i      instruction
	a      *int
	b      *int
	source string
}

type input struct {
	v []int
}

func (i *input) next() int {
	if len(i.v) == 0 {
		panic("no input")
	}
	n := i.v[0]
	i.v = i.v[1:]
	return n
}

func (alu *ALU) dump() {
	fmt.Println("DUMP:")
	fmt.Printf(" X:%d\n Y:%d\n z:%d\n w:%d", alu.X, alu.Y, alu.Z, alu.W)
	fmt.Println()
}

func (alu *ALU) GetValue(arg string) *int {
	switch arg {
	case "x":
		return &alu.X
	case "y":
		return &alu.Y
	case "z":
		return &alu.Z
	case "w":
		return &alu.W
	}

	v := util.TryParseInt(arg)
	return &v
}

var (
	instructions = map[string]instruction{
		"inp": func(a *int, b *int, i *input, c *int) {
			*a = i.next()
		},
		"add": func(a *int, b *int, i *input, c *int) {
			*a = *a + *b
		},
		"mul": func(a *int, b *int, i *input, c *int) {
			*a = *a * *b
		},
		"div": func(a *int, b *int, i *input, c *int) {
			*a = *a / *b
		},
		"mod": func(a *int, b *int, i *input, c *int) {
			*a = *a % *b
		},
		"eql": func(a *int, b *int, i *input, c *int) {
			if *a == *b {
				*a = 1
			} else {
				*a = 0
			}
		},
		"noop": func(a *int, b *int, i *input, c *int) {},
	}
)

func (alu *ALU) Compile(source []string) *Program {
	p := &Program{labels: make(map[string][]int)}
	lastLabel := ""
	for i, line := range source {

		if strings.Index(line, Label) == len(line)-1 {
			p.labels[line] = []int{i + 1, 0}
			if lastLabel != "" {
				p.labels[lastLabel][1] = i
			}
			lastLabel = line

			s := &statement{
				i:      instructions["noop"],
				a:      nil,
				b:      nil,
				source: line,
			}
			p.s = append(p.s, s)

			continue
		}

		parts := strings.Split(line, " ")

		i := instructions[parts[0]]

		var a, b *int
		a = alu.GetValue(parts[1])
		if len(parts) == 3 {
			b = alu.GetValue(parts[2])
		}

		s := &statement{
			i:      i,
			a:      a,
			b:      b,
			source: line,
		}

		p.s = append(p.s, s)
	}

	p.labels[lastLabel][1] = len(source) - 1

	return p
}

func (alu *ALU) Run(p *Program, inputs []int) {
	alu.reset()
	i := &input{v: inputs}
	c := 0
	for c < len(p.s) {
		s := p.s[c]
		s.i(s.a, s.b, i, &c)
		c++
	}

	if len(i.v) > 0 {
		panic("left over input")
	}
}
func (alu *ALU) RunLabel(p *Program, label string, x, y, z int, inputs []int) {
	alu.X = x
	alu.Y = y
	alu.Z = z
	alu.W = 0

	i := &input{v: inputs}

	c := p.labels[label][0]
	for c <= p.labels[label][1] {
		s := p.s[c]
		s.i(s.a, s.b, i, &c)
		c++
	}

	if len(i.v) > 0 {
		panic("left over inpu: label:" + label)
	}
}

func (alu *ALU) reset() {
	alu.X = 0
	alu.Y = 0
	alu.Z = 0
	alu.W = 0
}
