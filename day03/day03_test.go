package day03

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name        string
	diagnostics []string
	power       int
	lifeSupport int
}

func TestProcessDiagnostics(t *testing.T) {
	cases := []testCase{
		//gamma = 11011 (27), epsilon = 00100 (4) = 4*27 = 108
		{"Single report",
			[]string{"00100"},
			108,
			16,
		},
		{
			"Example",
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010"},
			198,
			230,
		},
		{
			"Real case",
			puzzleinput.ReadLines("03_input.txt"),
			3912944,
			4996233,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualPower, actualLifeSupport := ProcessDiagnostics(c.diagnostics)
			if actualPower != c.power {
				t.Errorf("expected power %d, got %d", c.power, actualPower)
			}
			if actualLifeSupport != c.lifeSupport {
				t.Errorf("expected life support %d, got %d", c.lifeSupport, actualLifeSupport)
			}
		})
	}
}
