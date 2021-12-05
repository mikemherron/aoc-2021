package day05

import (
	"AdventCode2021/util"
	"strings"
)

const PartsSeparator = "-> "

func direction(one, two int) int {
	if one > two {
		return -1
	} else if one < two {
		return 1
	} else {
		return 0
	}
}

func FindDangerousVents(vents []string, diagonals bool) int {
	total := 0
	cells := make([][]int, 0)
	for _, v := range vents {
		x1, y1, x2, y2 := parse(v)
		if !diagonals && y1 != y2 && x1 != x2 {
			continue
		}

		xC, yC := util.Max(x1, x2)+1, util.Max(y1, y2)+1
		if len(cells) < yC {
			cells = append(cells, make([][]int, yC-len(cells))...)
		}

		yD, xD := direction(y1, y2), direction(x1, x2)
		for x1 != x2+xD || y1 != y2+yD {
			if len(cells[y1]) < xC {
				cells[y1] = append(cells[y1], make([]int, xC-len(cells[y1]))...)
			}
			cells[y1][x1]++
			if cells[y1][x1] == 2 {
				total++
			}
			x1 += xD
			y1 += yD
		}
	}

	return total
}

func parse(v string) (int, int, int, int) {
	parts := strings.Split(v, PartsSeparator)
	start := util.SplitByCommaToInt(parts[0])
	end := util.SplitByCommaToInt(parts[1])

	return start[0], start[1], end[0], end[1]
}
