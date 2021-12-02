package day01

import (
	"AdventCode2020/puzzleinput"
	"testing"
)

type testCase struct {
	name              string
	window            int
	measurements      []int
	expectedIncreases int
}

func TestCountDepthIncreasesInWindow(t *testing.T) {

	cases := []*testCase{
		{"Single Measurement: Empty Input", 1, []int{}, 0},
		{"Single Measurement: Single Input", 1, []int{100}, 0},
		{"Single Measurement: Single increase", 1, []int{199, 200}, 1},
		{"Single Measurement: Example case", 1, []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 7},
		{"Single Measurement: Real case",1, puzzleinput.ReadIntsFrom("01_input.txt"), 1711},
		{"3 Window: Example case", 3, []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, 5},
		{"3 Window: Real case", 3, puzzleinput.ReadIntsFrom("01_input.txt"), 1743},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := countDepthIncreasesInWindow(c.measurements, c.window)
			if actual != c.expectedIncreases {
				t.Errorf("expected %d, got %d", c.expectedIncreases, actual)
			}
		})
	}
}

