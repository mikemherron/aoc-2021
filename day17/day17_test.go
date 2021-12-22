package day17

import (
	"testing"
)

func TestSimulateLauncher(t *testing.T) {

	type testCase struct {
		name  string
		input string
		maxY  int
		count int
	}

	cases := []testCase{
		{
			"Example",
			"target area: x=20..30, y=-10..-5",
			45,
			112,
		},
		{
			"Real",
			"target area: x=265..287, y=-103..-58",
			5253,
			1770,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualMaxY, actualCount := SimulateLauncher(c.input)
			if actualMaxY != c.maxY {
				t.Errorf("maxY %d, got %d", c.maxY, actualMaxY)
			}
			if actualCount != c.count {
				t.Errorf("count %d, got %d", c.count, actualCount)
			}
		})
	}
}
