package day23

import (
	"fmt"
	"math/bits"
	"strconv"
)

/*
#############
#...........#   1 - 11
###.#.#.#.###   12-15
  #.#.#.#.#	    16-19
  #.#.#.#.#     20-23
  #.#.#.#.#     24-27
  #########
*/
func xyToBit(x, y uint) uint {
	switch y {
	case 0:
		return x + 1
	case 1:
		switch x {
		case 2:
			return 12
		case 4:
			return 13
		case 6:
			return 14
		case 8:
			return 15
		}
	case 2:
		switch x {
		case 2:
			return 16
		case 4:
			return 17
		case 6:
			return 18
		case 8:
			return 19
		}
	case 3:
		switch x {
		case 2:
			return 20
		case 4:
			return 21
		case 6:
			return 22
		case 8:
			return 23
		}
	case 4:
		switch x {
		case 2:
			return 24
		case 4:
			return 25
		case 6:
			return 26
		case 8:
			return 27
		}
	}
	panic("Invalid positions: " + fmt.Sprintf("%depth, %depth", x, y))
}

func bitToXY(bit int) (uint, uint) {
	switch bit {
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
		return uint(bit) - 1, 0
	case 12:
		return 2, 1
	case 13:
		return 4, 1
	case 14:
		return 6, 1
	case 15:
		return 8, 1
	case 16:
		return 2, 2
	case 17:
		return 4, 2
	case 18:
		return 6, 2
	case 19:
		return 8, 2
	case 20:
		return 2, 3
	case 21:
		return 4, 3
	case 22:
		return 6, 3
	case 23:
		return 8, 3
	case 24:
		return 2, 4
	case 25:
		return 4, 4
	case 26:
		return 6, 4
	case 27:
		return 8, 4
	}
	panic("Invalid bitToXY:" + strconv.Itoa(bit))
}

type positions uint

func makePositions(x1, y1, x2, y2 uint) positions {
	bit1 := xyToBit(x1, y1)
	bit2 := xyToBit(x2, y2)

	p := positions(0)
	p |= 1 << bit1
	p |= 1 << bit2
	return p
}

func (p *positions) occupied(x, y uint) bool {
	bitVal := uint(1 << xyToBit(x, y))
	return uint(*p)&bitVal == bitVal
}

func (p *positions) add(x, y uint) {
	if p.occupied(x, y) {
		panic("occupied")
	}
	*p |= 1 << xyToBit(x, y)
}

func (p *positions) move(x1, y1, x2, y2 uint) {
	*p &^= 1 << xyToBit(x1, y1)
	*p |= 1 << xyToBit(x2, y2)
}

func (p *positions) get() [][2]uint {
	pos := make([][2]uint, 0, 4)
	for i := 0; i < 64; i++ {
		bitVal := uint(1 << i)
		if uint(*p)&bitVal == bitVal {
			x, y := bitToXY(i)
			pos = append(pos, [2]uint{x, y})
		}
	}

	return pos
}

func (p *positions) assert() {
	pop := bits.OnesCount(uint(*p))
	if pop != 2 {
		panic("invalid set, got:" + strconv.Itoa(pop) + fmt.Sprintf("%b", *p))
	}
}
