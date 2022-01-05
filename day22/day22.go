package day22

import (
	"AdventCode2021/util"
	"fmt"
	"regexp"
)

func PowerOn(input []string) int {
	power := 0
	cuboids := parseInstructions(input)
	for i, c := range cuboids {
		if c.on {
			cubesOn := c.size() - alreadyOn(c, cuboids[:i])
			power += cubesOn
		} else {
			cubesOff := alreadyOn(c, cuboids[:i])
			power -= cubesOff
		}
	}

	return power
}

func alreadyOn(c *Cuboid, cuboids []*Cuboid) int {
	on := 0
	for i := len(cuboids) - 1; i >= 0; i-- {
		o := cuboids[i]
		if o.intersects(c) {
			intersection := o.intersection(c)
			if o.on {
				on += intersection.size() - alreadyOn(intersection, cuboids[:i])
			} else {
				on -= alreadyOn(intersection, cuboids[:i])
			}
		}
	}

	return on
}

type Pos struct {
	x, y, z int
}

type Cuboid struct {
	from *Pos
	to   *Pos
	on   bool
}

func (c *Cuboid) String() string {
	return fmt.Sprintf("x=%d..%d, y=%d..%d, z=%d..%d", c.from.x, c.to.x, c.from.y, c.to.y, c.from.z, c.to.z)
}

func (c *Cuboid) intersection(o *Cuboid) *Cuboid {
	xFrom, xTo := axisIntersection(c.from.x, c.to.x, o.from.x, o.to.x)
	yFrom, yTo := axisIntersection(c.from.y, c.to.y, o.from.y, o.to.y)
	zFrom, zTo := axisIntersection(c.from.z, c.to.z, o.from.z, o.to.z)

	return &Cuboid{
		from: &Pos{xFrom, yFrom, zFrom},
		to:   &Pos{xTo, yTo, zTo},
	}
}

func (c *Cuboid) size() int {
	return (c.to.x - c.from.x + 1) * (c.to.y - c.from.y + 1) * (c.to.z - c.from.z + 1)
}

func (c Cuboid) intersects(o *Cuboid) bool {
	return axisIntersects(c.from.x, c.to.x, o.from.x, o.to.x) &&
		axisIntersects(c.from.y, c.to.y, o.from.y, o.to.y) &&
		axisIntersects(c.from.z, c.to.z, o.from.z, o.to.z)
}

func axisIntersects(a1, a2, b1, b2 int) bool {
	return (b1 >= a1 && b1 <= a2) || (b2 >= a1 && b2 <= a2) || (b1 <= a1 && b2 >= a2)
}

func axisIntersection(a1, a2, b1, b2 int) (int, int) {
	return util.Max(a1, b1), util.Min(a2, b2)
}

const (
	inputParts     = "((-*)\\d+)\\.\\.((-*)\\d+)"
	inputSeparator = ".."
	inputStateOn   = "on"
)

func parseInstructions(input []string) []*Cuboid {
	cubes := make([]*Cuboid, 0)
	for _, i := range input {
		state := i[:2]
		parts := regexp.MustCompile(inputParts).FindAllString(i, -1)

		xC := util.SplitToInt(parts[0], inputSeparator)
		yC := util.SplitToInt(parts[1], inputSeparator)
		zC := util.SplitToInt(parts[2], inputSeparator)

		s := Pos{xC[0], yC[0], zC[0]}
		e := Pos{xC[1], yC[1], zC[1]}

		cubes = append(cubes, &Cuboid{from: &s, to: &e, on: state == inputStateOn})
	}

	return cubes
}
