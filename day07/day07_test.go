package day07

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

type testCase struct {
	name    string
	input   []int
	fuel    int
	scaling bool
}

func TestAlignCrabs(t *testing.T) {

	cases := []testCase{
		{
			"Example",
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			37,
			false,
		},
		{
			"Real",
			puzzleinput.ReadCommaSeperatedInts("07_input.txt"),
			342730,
			false,
		},
		{
			"Example Scaling",
			[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			168,
			true,
		},
		{
			"Real Scaling",
			puzzleinput.ReadCommaSeperatedInts("07_input.txt"),
			0,
			true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := AlignCrabs(c.input, c.scaling)
			if actual != c.fuel {
				t.Errorf("expected %d, got %d", c.fuel, actual)
			}
		})
	}
}
