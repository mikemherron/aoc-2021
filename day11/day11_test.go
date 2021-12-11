package day11

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name      string
	input     []string
	flashes   int
	firstSync int
	steps     int
}

func TestCountFlashes(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			puzzleinput.ReadLines("11_example_input.txt"),
			1656,
			195,
			100,
		},
		{
			"Real",
			puzzleinput.ReadLines("11_input.txt"),
			1691,
			216,
			100,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualFlashes, actualFirstSync := CountFlashes(c.input, c.steps)
			if actualFlashes != c.flashes {
				t.Errorf("expected flashes %d, got %d", c.flashes, actualFlashes)
			}
			if actualFirstSync != c.firstSync {
				t.Errorf("expected firstSync %d, got %d", c.firstSync, actualFirstSync)
			}
		})
	}
}
