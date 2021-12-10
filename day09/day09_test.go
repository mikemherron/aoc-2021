package day09

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name   string
	input  []string
	risk   int
	basins int
}

func TestGetRiskLevel(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("09_example_input.txt"),
			15,
			1134,
		},
		{
			"Real Case",
			puzzleinput.ReadLines("09_input.txt"),
			562,
			1076922,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualRisk, actualBasin := GetRiskLevel(c.input)
			if actualRisk != c.risk {
				t.Errorf("expected risk %d, got %d", c.risk, actualRisk)
			}
			if actualBasin != c.basins {
				t.Errorf("expected basins %d, got %d", c.basins, actualBasin)
			}
		})
	}
}
