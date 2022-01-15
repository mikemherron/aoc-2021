package day23

import (
	"fmt"
	"math"
	"strings"
)

type amphipodType struct {
	cost  int
	xDest uint
}

var hallwayPositions = []uint{0, 1, 3, 5, 7, 9, 10}

var (
	Amber  = &amphipodType{1, 2}
	Bronze = &amphipodType{10, 4}
	Copper = &amphipodType{100, 6}
	Desert = &amphipodType{1000, 8}
)

type depth uint

var (
	DepthTwo  depth = 2
	DepthFour depth = 4
)

var (
	FinalAmberTwo = uint(1<<(xyToBit(Amber.xDest, 1)) |
		1<<(xyToBit(Amber.xDest, 2)))

	FinalAmberFour = FinalAmberTwo |
		1<<(xyToBit(Amber.xDest, 3)) |
		1<<(xyToBit(Amber.xDest, 4))

	FinalBronzeTwo = uint(1<<(xyToBit(Bronze.xDest, 1)) |
		1<<(xyToBit(Bronze.xDest, 2)))

	FinalBronzeFour = FinalBronzeTwo |
		1<<(xyToBit(Bronze.xDest, 3)) |
		1<<(xyToBit(Bronze.xDest, 4))

	FinalCopperTwo = uint(1<<(xyToBit(Copper.xDest, 1)) |
		1<<(xyToBit(Copper.xDest, 2)))

	FinalCopperFour = FinalCopperTwo |
		1<<(xyToBit(Copper.xDest, 3)) |
		1<<(xyToBit(Copper.xDest, 4))

	FinalDesertTwo = uint(1<<(xyToBit(Desert.xDest, 1)) |
		1<<(xyToBit(Desert.xDest, 2)))

	FinalDesertFour = FinalDesertTwo |
		1<<(xyToBit(Desert.xDest, 3)) |
		1<<(xyToBit(Desert.xDest, 4))
)

type typePositions struct {
	t *amphipodType
	p positions
}

func (t amphipodType) final(x, y uint) bool {
	return x == t.xDest && y > 0
}

type state struct {
	a    amphipods
	d    depth
	cost int
}

type amphipods struct {
	amber  typePositions
	bronze typePositions
	copper typePositions
	desert typePositions
}

type move struct {
	t            *amphipodType
	cost         int
	xFrom, yFrom uint
	xDest, yDest uint
}

func (s state) final() bool {
	if s.d == DepthFour {
		return uint(s.a.amber.p) == FinalAmberFour &&
			uint(s.a.bronze.p) == FinalBronzeFour &&
			uint(s.a.copper.p) == FinalCopperFour &&
			uint(s.a.desert.p) == FinalDesertFour
	} else {
		return uint(s.a.amber.p) == FinalAmberTwo &&
			uint(s.a.bronze.p) == FinalBronzeTwo &&
			uint(s.a.copper.p) == FinalCopperTwo &&
			uint(s.a.desert.p) == FinalDesertTwo
	}
}

func (s state) moves() []*move {
	moves := make([]*move, 0, 16)
	moves = s.movesForSet(s.a.amber, moves)
	moves = s.movesForSet(s.a.bronze, moves)
	moves = s.movesForSet(s.a.copper, moves)
	moves = s.movesForSet(s.a.desert, moves)

	return moves
}

func (s state) movesForSet(as typePositions, moves []*move) []*move {
	for _, pos := range as.p.get() {

		x, y := pos[0], pos[1]
		if s.lockedIn(x, y, as.t) {
			continue
		}

		if a, yDest := s.roomAvailable(as.t); a {
			if ok, distance := s.path(x, y, as.t.xDest, yDest); ok {
				moves = append(moves, &move{as.t, distance * as.t.cost, x, y, as.t.xDest, yDest})
			}
		}

		if y == 0 {
			continue
		}

		for _, xDest := range hallwayPositions {
			if ok, distance := s.path(x, y, xDest, 0); ok {
				moves = append(moves, &move{as.t, distance * as.t.cost, x, y, xDest, 0})
			}
		}
	}

	return moves
}

func (s state) path(x1, y1, x2, y2 uint) (bool, int) {
	moves := 0

	//If in a room space, move up to corridor
	for y1 > 0 {
		moves++
		y1--
		if o, _ := s.occupied(x1, y1); o {
			return false, 0
		}
	}

	//Move along the corridor
	dir := direction(x1, x2)
	for x1 != x2 {
		moves++
		x1 = uint(int(x1) + dir)
		if o, _ := s.occupied(x1, y1); o {
			return false, 0
		}
	}

	//Move down in to room
	for y1 < y2 {
		moves++
		y1++
		if o, _ := s.occupied(x1, y1); o {
			return false, 0
		}
	}

	return true, moves
}

func (s state) occupied(x, y uint) (bool, *amphipodType) {
	if s.a.amber.p.occupied(x, y) {
		return true, Amber
	}
	if s.a.bronze.p.occupied(x, y) {
		return true, Bronze
	}
	if s.a.copper.p.occupied(x, y) {
		return true, Copper
	}
	if s.a.desert.p.occupied(x, y) {
		return true, Desert
	}

	return false, Amber
}

func (s state) lockedIn(x, y uint, t *amphipodType) bool {
	if x != t.xDest || y < 1 {
		return false
	}

	for i := y; i <= uint(s.d); i++ {
		o, ot := s.occupied(x, i)
		if !o || ot != t {
			return false
		}
	}

	return true
}

func (s state) roomAvailable(t *amphipodType) (bool, uint) {
	var next uint
	for y := uint(1); y <= uint(s.d); y++ {
		o, ot := s.occupied(t.xDest, y)
		if !o {
			next = y
		}
		if o && ot != t {
			return false, 0
		}
	}

	return true, next
}

func (s state) apply(m *move) state {
	moved := state{
		a:    s.a,
		d:    s.d,
		cost: s.cost + m.cost,
	}

	if m.t == Amber {
		moved.a.amber.p.move(m.xFrom, m.yFrom, m.xDest, m.yDest)
	} else if m.t == Bronze {
		moved.a.bronze.p.move(m.xFrom, m.yFrom, m.xDest, m.yDest)
	} else if m.t == Copper {
		moved.a.copper.p.move(m.xFrom, m.yFrom, m.xDest, m.yDest)
	} else if m.t == Desert {
		moved.a.desert.p.move(m.xFrom, m.yFrom, m.xDest, m.yDest)
	}

	return moved
}

func (s state) Print() {

	char := func(x, y uint) string {
		if s.a.amber.p.occupied(x, y) {
			return "A"
		} else if s.a.bronze.p.occupied(x, y) {
			return "B"
		} else if s.a.copper.p.occupied(x, y) {
			return "C"
		} else if s.a.desert.p.occupied(x, y) {
			return "D"
		}

		return "."
	}

	fmt.Println("#############")
	fmt.Printf("#")
	for i := uint(0); i <= 10; i++ {
		fmt.Printf(char(i, 0))
	}
	fmt.Printf("#\n")
	fmt.Printf("###%s#%s#%s#%s###\n", char(2, 1), char(4, 1), char(6, 1), char(8, 1))
	fmt.Printf("  #%s#%s#%s#%s#  \n", char(2, 2), char(4, 2), char(6, 2), char(8, 2))
	fmt.Printf("  #%s#%s#%s#%s#  \n", char(2, 3), char(4, 3), char(6, 3), char(8, 3))
	fmt.Printf("  #%s#%s#%s#%s#  \n", char(2, 4), char(4, 4), char(6, 4), char(8, 4))
	fmt.Println("  #########")
}

func (s state) equals(e state) bool {
	return s.a == e.a
}

func LeastEnergy(input string) int {
	possible, min := leastEnergyFrom(parseState(input), make(map[amphipods][]*move), math.MaxInt)
	if !possible {
		panic("not possible")
	}

	return min
}

func leastEnergyFrom(s state, cache map[amphipods][]*move, min int) (bool, int) {
	if s.cost >= min {
		return false, 0
	}

	if s.final() {
		return true, s.cost
	}

	var moves []*move
	if _, ok := cache[s.a]; ok {
		moves = cache[s.a]
	} else {
		moves = s.moves()
		cache[s.a] = moves
	}

	if len(moves) == 0 {
		return false, 0
	}

	hasValidMoves := false
	for _, m := range moves {
		possible, cost := leastEnergyFrom(s.apply(m), cache, min)
		if possible {
			hasValidMoves = true
			min = cost
		}
	}

	if !hasValidMoves {
		return false, 0
	}

	return true, min
}

func parseState(input string) state {

	//		#############
	//		#...........#
	//		###B#C#B#D###
	//		  #D#C#B#A#
	//		  #D#B#A#C#
	//		  #A#D#C#A#
	//		  #########

	s := state{}
	s.a.amber.t = Amber
	s.a.bronze.t = Bronze
	s.a.copper.t = Copper
	s.a.desert.t = Desert

	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 7 {
		s.d = DepthFour
	} else if len(lines) == 5 {
		s.d = DepthTwo
	} else {
		panic("Unsupported depth")
	}

	for x := 1; x < 12; x++ {
		if lines[1][x] != '.' {
			switch lines[1][x] {
			case 'A':
				s.a.amber.p.add(uint(x-1), 0)
			case 'B':
				s.a.bronze.p.add(uint(x-1), 0)
			case 'C':
				s.a.copper.p.add(uint(x-1), 0)
			case 'D':
				s.a.desert.p.add(uint(x-1), 0)
			}
		}
	}

	for x := 3; x <= 9; x += 2 {
		for y := 2; y < 2+int(s.d); y++ {
			if lines[y][x] != '.' {
				switch lines[y][x] {
				case 'A':
					s.a.amber.p.add(uint(x-1), uint(y-1))
				case 'B':
					s.a.bronze.p.add(uint(x-1), uint(y-1))
				case 'C':
					s.a.copper.p.add(uint(x-1), uint(y-1))
				case 'D':
					s.a.desert.p.add(uint(x-1), uint(y-1))
				}
			}
		}
	}
	return s
}

func direction(x1, x2 uint) int {
	if x1 > x2 {
		return -1
	}

	return 1
}
