package day21

import (
	"AdventCode2021/util"
	"fmt"
)

func PlayQuantum(p1, p2 int) int {
	starting := state{
		active: player{num: 1, pos: p1},
		next:   player{num: 2, pos: p2},
	}

	results := play(starting, make(map[state]map[int]int))

	return util.Max(results[1], results[2])
}

type player struct {
	num   int
	pos   int
	score int
}

func (p *player) update(roll int) {
	p.pos += roll
	if p.pos > 10 {
		p.pos = p.pos % 10
		if p.pos == 0 {
			p.pos = 10
		}
	}

	p.score += p.pos
}

type state struct {
	active player
	next   player
}

func play(s state, outcomes map[state]map[int]int) map[int]int {
	if cached, ok := outcomes[s]; ok {
		return cached
	}

	wins := map[int]int{
		1: 0,
		2: 0,
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				active := s.active
				active.update(i + j + k)
				if active.score >= 21 {
					wins[active.num]++
					continue
				}

				results := play(state{active: s.next, next: active}, outcomes)
				wins[1] += results[1]
				wins[2] += results[2]
			}
		}
	}

	outcomes[s] = wins

	return wins
}

type dice struct {
	value int
	rolls int
}

func (d *dice) roll() int {
	r := 0
	for i := 0; i < 3; i++ {
		r += d.value
		d.value++
		d.rolls++

		if d.value > 100 {
			d.value = 1
		}
	}

	fmt.Println()

	return r
}

func PlayDeterministic(p1, p2 int) int {

	d := &dice{value: 1, rolls: 0}

	players := []*player{
		{num: 1, pos: p1, score: 0},
		{num: 2, pos: p2, score: 0},
	}

	lastScore := 0
	for {
		for _, p := range players {
			p.update(d.roll())
			if p.score >= 1000 {
				return d.rolls * lastScore
			}

			lastScore = p.score
		}
	}
}
