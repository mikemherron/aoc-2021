package day17

import (
	"AdventCode2021/util"
	"math"
	"regexp"
)

const InputPattern = "((-*)\\d+)"

type result struct {
	success bool
	yMax    int
}

func SimulateLauncher(input string) (int, int) {
	maxY, count, target := math.MinInt, 0, parseTarget(input)
	for x := 1; x <= target[1][0]; x++ {
		for y := target[1][1]; y < 1000; y++ {
			r := run(x, y, target)
			if r.success {
				count++
				if r.yMax > maxY {
					maxY = r.yMax
				}
			}
		}
	}

	return maxY, count
}

func run(vX, vY int, target [][]int) result {
	x, y, r := 0, 0, result{yMax: math.MinInt}
	for {
		x += vX
		y += vY
		if y > r.yMax {
			r.yMax = y
		}

		vY--
		if vX > 0 {
			vX--
		} else if vX < 0 {
			vX++
		}

		if x >= target[0][0] && x <= target[1][0] && y <= target[0][1] && y >= target[1][1] {
			r.success = true
			break
		}

		if x > target[1][0] || y < target[1][1] {
			break
		}
	}

	return r
}

func parseTarget(input string) [][]int {
	parts := util.AsInts(regexp.MustCompile(InputPattern).FindAllString(input, -1))
	return [][]int{{parts[0], parts[3]}, {parts[1], parts[2]}}
}
