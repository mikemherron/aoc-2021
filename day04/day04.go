package day04

import (
	"AdventCode2021/util"
	"strings"
)

type board struct {
	nums   map[int][]int
	rows   []int
	cols   []int
	total  int
	active bool
}

func (b *board) add(n, r, c int) {
	b.nums[n] = make([]int, 2)
	b.nums[n][0] = r
	b.nums[n][1] = c
	b.total += n
}

func (b *board) mark(n int) bool {
	b.total -= n
	b.rows[b.nums[n][0]]++
	b.cols[b.nums[n][1]]++
	return b.rows[b.nums[n][0]] == 5 || b.cols[b.nums[n][1]] == 5
}

func newBoard() *board {
	return &board{
		nums:   make(map[int][]int),
		rows:   make([]int, 5),
		cols:   make([]int, 5),
		active: true,
	}
}

func Bingo(s []string) (int, int) {

	nums := make([]int, 0)
	for _, val := range strings.Split(s[0], ",") {
		nums = append(nums, util.TryParseInt(val))
	}

	boardIdx := make(map[int][]*board)
	boardsInput := s[1:]

	for i := 0; i < len(boardsInput); i += 6 {
		b := newBoard()
		for rNum, r := range boardsInput[i+1 : i+6] {
			cols := strings.Split(r, " ")
			cNum := 0
			for _, c := range cols {
				if c == "" {
					continue
				}
				num := util.TryParseInt(c)
				b.add(num, rNum, cNum)
				if _, ok := boardIdx[num]; !ok {
					boardIdx[num] = make([]*board, 0)
				}
				boardIdx[num] = append(boardIdx[num], b)
				cNum++
			}
		}
	}

	first, last := 0, 0
	for _, i := range nums {
		for _, b := range boardIdx[i] {
			if b.active && b.mark(i) {
				if first == 0 {
					first = b.total * i
				}
				last = b.total * i
				b.active = false
			}
		}
	}

	return first, last
}
