package alu

import (
	"testing"
)

func TestALU(t *testing.T) {

	type testCase struct {
		name     string
		program  []string
		input    []int
		expected map[string]int
	}

	cases := []testCase{
		{
			name: "Negate number",
			program: []string{
				"inp x",
				"mul x -1",
			},
			input: []int{
				10,
			},
			expected: map[string]int{
				"x": -10,
			},
		},
		{
			name: "3 times as big true",
			program: []string{
				"inp z",
				"inp x",
				"mul z 3",
				"eql z x",
			},
			input: []int{
				5, 15,
			},
			expected: map[string]int{
				"z": 1,
			},
		},
		{
			name: "3 times as big false",
			program: []string{
				"inp z",
				"inp x",
				"mul z 3",
				"eql z x",
			},
			input: []int{
				1, 15,
			},
			expected: map[string]int{
				"z": 0,
			},
		},
	}

	for _, c := range cases {
		alu := NewALU()
		p := alu.Compile(c.program)
		alu.Run(p, c.input)
		for r, e := range c.expected {
			if *(alu.GetValue(r)) != e {
				t.Errorf("Expected value %d for register %source, got %d", e, r, *(alu.GetValue(r)))
			}
		}
	}

}
