package day19

import (
	"AdventCode2021/util"
	"fmt"
	"math"
)

func CountBeacons(input []string, threshold int) (int, int) {

	scanners := parse(input)

	all := scanners[0]

	list := make([]*scanner, len(scanners)-1)
	copy(list, scanners[1:])
	i := 0
	for len(list) > 0 {
		if findCommonBeacons(list[i], threshold, all) {
			list = append(list[:i], list[i+1:]...)
			i = 0
		} else {
			i++
			if i > len(list) {
				panic("Ran out of tries")
			}
		}
	}

	max := math.MinInt
	for i, s := range scanners {
		for j, s2 := range scanners {
			dist := s.distance(s2)
			if dist > max {
				fmt.Printf("Max between %d and %d: %d\n", i, j, dist)
				max = dist
			}
		}
	}
	return len(all.beacons), max
}

func findCommonBeacons(scanner *scanner, threshold int, all *scanner) bool {
	for _, set := range scanner.permutationSets {
		for beacon := range set {
			for knownBeacon := range all.beacons {

				xOffset := knownBeacon.x - beacon.x
				yOffset := knownBeacon.y - beacon.y
				zOffset := knownBeacon.z - beacon.z

				offsets := offsetBeacons(set, xOffset, yOffset, zOffset)
				matches := all.matches(offsets)

				if matches >= threshold {
					scanner.pos.x = -xOffset
					scanner.pos.y = -yOffset
					scanner.pos.z = -zOffset
					for offset := range offsets {
						all.beacons[offset] = true
					}

					return true
				}
			}
		}
	}

	return false
}

type position struct {
	x, y, z int
}

func (p position) string() string {
	return fmt.Sprintf("x:%d,y:%d,z:%d", p.x, p.y, p.z)
}

func (s scanner) distance(s2 *scanner) int {
	return util.Abs(s.pos.x-s2.pos.x) + (s.pos.y - s2.pos.y) + (s.pos.z - s2.pos.z)
}

func newScanner() *scanner {
	return &scanner{beacons: make(map[position]bool, 0)}
}

type scanner struct {
	pos             position
	beacons         map[position]bool
	permutationSets []map[position]bool
}

func (s scanner) matches(beacons map[position]bool) int {
	match := 0
	for b := range beacons {
		if s.beacons[b] {
			match++
		}
	}

	return match
}

func offsetBeacons(b map[position]bool, x, y, z int) map[position]bool {
	offset := make(map[position]bool, 0)
	for b := range b {
		offsetBeacon := b
		offsetBeacon.x += x
		offsetBeacon.y += y
		offsetBeacon.z += z

		offset[offsetBeacon] = true
	}

	return offset
}

func parse(input []string) []*scanner {
	scanners := make([]*scanner, 0)
	var current *scanner
	for _, s := range input {
		if s == "" {
			continue
		}

		if string(s[0:3]) == "---" {
			if current != nil {
				current.makePermutations()
			}
			current = newScanner()
			scanners = append(scanners, current)
		} else {
			coords := util.SplitByCommaToInt(s)
			current.beacons[position{
				x: coords[0],
				y: coords[1],
				z: coords[2],
			}] = true
		}
	}

	current.makePermutations()

	return scanners
}

func (s *scanner) makePermutations() {
	s.permutationSets = make([]map[position]bool, 48)
	for i := 0; i < 48; i++ {
		s.permutationSets[i] = make(map[position]bool, len(s.beacons))
	}

	for b := range s.beacons {
		//I'm sorry. I am so, so sorry.
		s.permutationSets[0][position{x: -b.x, y: -b.y, z: -b.z}] = true
		s.permutationSets[1][position{x: -b.x, y: -b.y, z: b.z}] = true
		s.permutationSets[2][position{x: -b.x, y: b.y, z: -b.z}] = true
		s.permutationSets[3][position{x: -b.x, y: b.y, z: b.z}] = true
		s.permutationSets[4][position{x: b.x, y: -b.y, z: -b.z}] = true
		s.permutationSets[5][position{x: b.x, y: -b.y, z: b.z}] = true
		s.permutationSets[6][position{x: b.x, y: b.y, z: -b.z}] = true
		s.permutationSets[7][position{x: b.x, y: b.y, z: b.z}] = true

		s.permutationSets[8][position{x: -b.x, y: -b.z, z: -b.y}] = true
		s.permutationSets[9][position{x: -b.x, y: -b.z, z: b.y}] = true
		s.permutationSets[10][position{x: -b.x, y: b.z, z: -b.y}] = true
		s.permutationSets[11][position{x: -b.x, y: b.z, z: b.y}] = true
		s.permutationSets[12][position{x: b.x, y: -b.z, z: -b.y}] = true
		s.permutationSets[13][position{x: b.x, y: -b.z, z: b.y}] = true
		s.permutationSets[14][position{x: b.x, y: b.z, z: -b.y}] = true
		s.permutationSets[15][position{x: b.x, y: b.z, z: b.y}] = true

		s.permutationSets[16][position{x: -b.y, y: -b.x, z: -b.z}] = true
		s.permutationSets[17][position{x: -b.y, y: -b.x, z: b.z}] = true
		s.permutationSets[18][position{x: -b.y, y: b.x, z: -b.z}] = true
		s.permutationSets[19][position{x: -b.y, y: b.x, z: b.z}] = true
		s.permutationSets[20][position{x: b.y, y: -b.x, z: -b.z}] = true
		s.permutationSets[21][position{x: b.y, y: -b.x, z: b.z}] = true
		s.permutationSets[22][position{x: b.y, y: b.x, z: -b.z}] = true
		s.permutationSets[23][position{x: b.y, y: b.x, z: b.z}] = true

		s.permutationSets[24][position{x: -b.y, y: -b.z, z: -b.x}] = true
		s.permutationSets[25][position{x: -b.y, y: -b.z, z: b.x}] = true
		s.permutationSets[26][position{x: -b.y, y: b.z, z: -b.x}] = true
		s.permutationSets[27][position{x: -b.y, y: b.z, z: b.x}] = true
		s.permutationSets[28][position{x: b.y, y: -b.z, z: -b.x}] = true
		s.permutationSets[29][position{x: b.y, y: -b.z, z: b.x}] = true
		s.permutationSets[30][position{x: b.y, y: b.z, z: -b.x}] = true
		s.permutationSets[31][position{x: b.y, y: b.z, z: b.x}] = true

		s.permutationSets[32][position{x: -b.z, y: -b.x, z: -b.y}] = true
		s.permutationSets[33][position{x: -b.z, y: -b.x, z: b.y}] = true
		s.permutationSets[34][position{x: -b.z, y: b.x, z: -b.y}] = true
		s.permutationSets[35][position{x: -b.z, y: b.x, z: b.y}] = true
		s.permutationSets[36][position{x: b.z, y: -b.x, z: -b.y}] = true
		s.permutationSets[37][position{x: b.z, y: -b.x, z: b.y}] = true
		s.permutationSets[38][position{x: b.z, y: b.x, z: -b.y}] = true
		s.permutationSets[39][position{x: b.z, y: b.x, z: b.y}] = true

		s.permutationSets[40][position{x: -b.z, y: -b.y, z: -b.x}] = true
		s.permutationSets[41][position{x: -b.z, y: -b.y, z: b.x}] = true
		s.permutationSets[42][position{x: -b.z, y: b.y, z: -b.x}] = true
		s.permutationSets[43][position{x: -b.z, y: b.y, z: b.x}] = true
		s.permutationSets[44][position{x: b.z, y: -b.y, z: -b.x}] = true
		s.permutationSets[45][position{x: b.z, y: -b.y, z: b.x}] = true
		s.permutationSets[46][position{x: b.z, y: b.y, z: -b.x}] = true
		s.permutationSets[47][position{x: b.z, y: b.y, z: b.x}] = true

	}

}
