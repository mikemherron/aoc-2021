package day02

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name         string
	instructions []string
	expected     int
}

func TestFinalSubPosition(t *testing.T) {
	runSubTests(t, RunBasicInstructions, []*testCase{
		{
			"Move once each way",
			[]string{"forward 1", "down 1"},
			1,
		},
		{
			"Example",
			[]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			150,
		},
		{
			"Real case",
			puzzleinput.ReadLinesFrom("02_input.txt"),
			1484118,
		},
	})
}

func TestFinalSubPositionWithAim(t *testing.T) {
	runSubTests(t, RunAimInstructions, []*testCase{
		{
			"Example",
			[]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			900,
		},
		{
			"Real case",
			puzzleinput.ReadLinesFrom("02_input.txt"),
			1463827010,
		},
	})
}

func runSubTests(t *testing.T, run func([]string) int, cases []*testCase) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := run(c.instructions)
			if actual != c.expected {
				t.Errorf("expected %d, got %d", c.expected, actual)
			}
		})
	}
}
