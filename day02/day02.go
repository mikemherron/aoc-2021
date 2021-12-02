package day02

import (
	"AdventCode2020/util"
	"strings"
)

const (
	Down    = "down"
	Up      = "up"
	Forward = "forward"
)

var (
	basicCommands = map[string]command{
		Forward: func(s *sub, i int) {
			s.pos += i
		},
		Down: func(s *sub, i int) {
			s.depth += i
		},
		Up: func(s *sub, i int) {
			s.depth -= i
		},
	}

	aimCommands = map[string]command{
		Forward: func(s *sub, i int) {
			s.pos += i
			s.depth += s.aim * i
		},
		Down: func(s *sub, i int) {
			s.aim += i
		},
		Up: func(s *sub, i int) {
			s.aim -= i
		},
	}
)

func RunBasicInstructions(instructions []string) int {
	s := sub{}
	s.run(instructions, basicCommands)
	return s.final()
}

func RunAimInstructions(instructions []string) int {
	s := sub{}
	s.run(instructions, aimCommands)
	return s.final()
}

type sub struct {
	pos   int
	depth int
	aim   int
}

type command func(*sub, int)

func (s *sub) final() int {
	return s.pos * s.depth
}

func (s *sub) run(instructions []string, commands map[string]command) {
	for _, i := range instructions {
		parts := strings.Split(i, " ")
		direction := parts[0]
		if cmd, ok := commands[direction]; ok {
			cmd(s, util.TryParseInt(parts[1]))
		}
	}
}
