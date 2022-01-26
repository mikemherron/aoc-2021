package day19

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

func TestCountBeacons(t *testing.T) {

	type testCases struct {
		name        string
		input       []string
		threshold   int
		beacons     int
		maxDistance int
	}

	cases := []testCases{
		{
			"Example",
			puzzleinput.ReadLines("19_example.txt"),
			12,
			79,
			3621,
		},
		{
			"Real",
			puzzleinput.ReadLines("19_input.txt"),
			12,
			419,
			13210,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualBeacons, actualDistance := CountBeacons(c.input, c.threshold)
			if actualBeacons != c.beacons {
				t.Errorf("expected beacons %d, got %d", c.beacons, actualBeacons)
			}
			if actualDistance != c.maxDistance {
				t.Errorf("expected max dist %d, got %d", c.maxDistance, actualDistance)
			}
		})
	}
}

func TestMakePermuations(t *testing.T) {

	input := newScanner()
	input.beacons[position{x: 1, y: 2, z: 3}] = true
	input.makePermutations()

	expected := []position{
		{x: -1, y: -2, z: -3},
		{x: -1, y: -2, z: 3},
		{x: -1, y: 2, z: -3},
		{x: -1, y: 2, z: 3},
		{x: 1, y: -2, z: 3},
		{x: 1, y: 2, z: -3},
		{x: 1, y: 2, z: 3},
		{x: -1, y: -3, z: -2},
		{x: -1, y: -3, z: 2},
		{x: -1, y: 3, z: -2},
		{x: -1, y: 3, z: 2},
		{x: 1, y: -3, z: 2},
		{x: 1, y: 3, z: -2},
		{x: 1, y: 3, z: 2},
		{x: -2, y: -1, z: -3},
		{x: -2, y: -1, z: 3},
		{x: -2, y: 1, z: -3},
		{x: -2, y: 1, z: 3},
		{x: 2, y: -1, z: -3},
		{x: 2, y: -1, z: 3},
		{x: 2, y: 1, z: -3},
		{x: 2, y: 1, z: 3},
		{x: -2, y: -3, z: -1},
		{x: -2, y: -3, z: 1},
		{x: -2, y: 3, z: -1},
		{x: -2, y: 3, z: 1},
		{x: 2, y: -3, z: -1},
		{x: 2, y: -3, z: 1},
		{x: 2, y: 3, z: -1},
		{x: 2, y: 3, z: 1},
		{x: -3, y: -1, z: -2},
		{x: -3, y: -1, z: 2},
		{x: -3, y: 1, z: -2},
		{x: -3, y: -1, z: 2},
		{x: -3, y: 1, z: -2},
		{x: -3, y: 1, z: 2},
		{x: 3, y: -1, z: -2},
		{x: 3, y: -1, z: 2},
		{x: 3, y: 1, z: -2},
		{x: 3, y: 1, z: 2},
		{x: -3, y: -2, z: -1},
		{x: -3, y: -2, z: 1},
		{x: -3, y: -2, z: 1},
		{x: -3, y: 2, z: -1},
		{x: -3, y: 2, z: 1},
		{x: 3, y: -2, z: -1},
		{x: 3, y: -2, z: 1},
		{x: 3, y: 2, z: -1},
		{x: 3, y: 2, z: 1},
	}

	for _, e := range expected {
		found := false
		for _, p := range input.permutationSets {
			if p[e] {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("did not find permuation %s", e.string())
		}
	}
}
