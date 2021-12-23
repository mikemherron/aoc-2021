package day20

import (
	"AdventCode2021/puzzleinput"
	"testing"
)

func TestEnhancedPixels(t *testing.T) {

	type testCases struct {
		name   string
		input  []string
		passes int
		pixels int
	}

	cases := []testCases{
		{
			"Example",
			puzzleinput.ReadLines("20_example.txt"),
			2,
			35,
		},
		{
			"Real",
			puzzleinput.ReadLines("20_input.txt"),
			2,
			5597,
		},
		{
			"Example 50 passes",
			puzzleinput.ReadLines("20_example.txt"),
			50,
			3351,
		},
		{
			"Real 50 passes",
			puzzleinput.ReadLines("20_input.txt"),
			50,
			18723,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actualPixels := EnhancedPixels(c.input, c.passes)
			if actualPixels != c.pixels {
				t.Errorf("expected pixels %d, got %d", c.pixels, actualPixels)
			}
		})
	}
}
