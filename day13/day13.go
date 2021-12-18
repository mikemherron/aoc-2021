package day13

import (
	"AdventCode2021/util"
	"fmt"
	"strings"
)

type point struct {
	x, y int
}

func CountDots(input []string) int {
	dots := make(map[point]bool)

	i := 0
	for {
		if strings.TrimSpace(input[i]) == "" {
			i++
			break
		}

		points := util.SplitByCommaToInt(input[i])
		dots[point{points[0], points[1]}] = true
		i++
	}

	for i < len(input) {
		firstFoldParts := strings.Split(input[i], " ")
		coords := firstFoldParts[2]
		coordParts := strings.Split(coords, "=")
		dir := coordParts[0]
		size := util.TryParseInt(coordParts[1])

		if dir == "y" {
			for k := range dots {
				if k.y > size && dots[k] {
					dots[k] = false
					diff := k.y - size
					newY := size - diff
					dots[point{k.x, newY}] = true

				}
			}
		} else {
			for k := range dots {
				if k.x > size && dots[k] {
					dots[k] = false
					diff := k.x - size
					newX := size - diff
					dots[point{newX, k.y}] = true

				}
			}
		}
		i++
	}

	t, maxX, maxY := 0, 0, 0
	for k, v := range dots {
		if v {
			t++
			if k.x > maxX {
				maxX = k.x
			}
			if k.y > maxY {
				maxY = k.y
			}
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	y := 0
	for y <= maxY {
		x := 0
		for x <= maxX {
			c := "."
			if dots[point{x, y}] {
				c = "#"
			}
			fmt.Print(c)
			x++
		}
		fmt.Print("\n")
		y++
	}
	fmt.Println()

	return t
}
