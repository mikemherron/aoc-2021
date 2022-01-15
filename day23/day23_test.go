package day23

import (
	"testing"
)

func TestIsFinal(t *testing.T) {
	s := parseState("" +
		"#############\n" +
		"#...........#\n" +
		"###A#B#C#D###\n" +
		"  #A#B#C#D#  \n" +
		"  #A#B#C#D#  \n" +
		"  #A#B#C#D#  \n" +
		"  #########  \n",
	)

	if !s.final() {
		s.Print()
		t.Errorf("State should be final")
	}
}

func TestLockedIn(t *testing.T) {

	type position struct {
		t    *amphipodType
		x, y uint
	}

	type testCase struct {
		s        string
		lockedIn map[position]bool
	}

	test := testCase{
		s: "" +
			"#############\n" +
			"#...D.B.....#\n" +
			"###A#B#C#.###\n" +
			"  #A#D#C#.#  \n" +
			"  #A#B#C#D#  \n" +
			"  #A#B#C#D#  \n" +
			"  #########  \n",

		lockedIn: map[position]bool{
			position{Amber, 2, 1}: true,
			position{Amber, 2, 2}: true,
			position{Amber, 2, 3}: true,
			position{Amber, 2, 4}: true,

			position{Desert, 3, 0}: false,

			position{Bronze, 4, 1}: false,
			position{Desert, 4, 2}: false,
			position{Bronze, 4, 3}: true,
			position{Bronze, 4, 4}: true,

			position{Copper, 6, 1}: true,
			position{Copper, 6, 2}: true,
			position{Copper, 6, 3}: true,
			position{Copper, 6, 4}: true,

			position{Desert, 8, 3}: true,
			position{Desert, 8, 4}: true,
		},
	}

	s := parseState(test.s)
	for k, expected := range test.lockedIn {
		lockedIn := s.lockedIn(k.x, k.y, k.t)
		if lockedIn != expected {
			t.Errorf("Expected %depth,%depth of type %v to be %v but was %v", k.x, k.y, k.t, expected, lockedIn)
		}
	}
}

func TestRoomAvailable(t *testing.T) {

	type expected struct {
		available bool
		y         uint
	}

	type testCase struct {
		s         string
		available map[*amphipodType]expected
	}

	test := testCase{
		s: "" +
			"#############\n" +
			"#.C.D.B....B#\n" +
			"###A#.#.#.###\n" +
			"  #A#D#C#.#  \n" +
			"  #A#B#C#D#  \n" +
			"  #A#B#C#D#  \n" +
			"  #########  \n",

		available: map[*amphipodType]expected{
			Amber:  {true, 0},
			Bronze: {false, 0},
			Copper: {true, 1},
			Desert: {true, 2},
		},
	}

	s := parseState(test.s)
	for k, e := range test.available {
		available, y := s.roomAvailable(k)
		if available != e.available {
			t.Errorf("Expected type %v to be available %v but was %v", *k, e.available, available)
		}

		if available {
			if y != e.y {
				t.Errorf("Expected type %v yFrom to be %depth was %v", k, e.y, y)
			}
		}
	}
}

func TestMoves(t *testing.T) {

	type testCase struct {
		start      string
		expectMove string
	}

	cases := []testCase{
		{
			start: "" +
				"#############\n" +
				"#...........#\n" +
				"###B#C#B#D###\n" +
				"  #D#C#B#A#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#..........D#\n" +
				"###B#C#B#.###\n" +
				"  #D#C#B#A#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
		{
			start: "" +
				"#############\n" +
				"#..........D#\n" +
				"###B#C#B#.###\n" +
				"  #D#C#B#A#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#A.........D#\n" +
				"###B#C#B#.###\n" +
				"  #D#C#B#.#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
		{
			start: "" +
				"#############\n" +
				"#A.........D#\n" +
				"###B#C#B#.###\n" +
				"  #D#C#B#.#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#A........BD#\n" +
				"###B#C#.#.###\n" +
				"  #D#C#B#.#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
		{
			start: "" +
				"#############\n" +
				"#A........BD#\n" +
				"###B#C#.#.###\n" +
				"  #D#C#B#.#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#A......B.BD#\n" +
				"###B#C#.#.###\n" +
				"  #D#C#.#.#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
		{
			start: "" +
				"#############\n" +
				"#A......B.BD#\n" +
				"###B#C#.#.###\n" +
				"  #D#C#.#.#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#AA.....B.BD#\n" +
				"###B#C#.#.###\n" +
				"  #D#C#.#.#  \n" +
				"  #D#B#.#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
		{
			start: "" +
				"#############\n" +
				"#AA.....B.BD#\n" +
				"###B#C#.#.###\n" +
				"  #D#C#.#.#  \n" +
				"  #D#B#.#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#AA.....B.BD#\n" +
				"###B#.#.#.###\n" +
				"  #D#C#.#.#  \n" +
				"  #D#B#C#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
		{
			start: "" +
				"#############\n" +
				"#AA.....B.BD#\n" +
				"###B#.#.#.###\n" +
				"  #D#C#.#.#  \n" +
				"  #D#B#C#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#AA.....B.BD#\n" +
				"###B#.#.#.###\n" +
				"  #D#.#C#.#  \n" +
				"  #D#B#C#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
		{
			start: "" +
				"#############\n" +
				"#AA.....B.BD#\n" +
				"###B#.#.#.###\n" +
				"  #D#.#C#.#  \n" +
				"  #D#B#C#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			expectMove: "" +
				"#############\n" +
				"#AA...B.B.BD#\n" +
				"###B#.#.#.###\n" +
				"  #D#.#C#.#  \n" +
				"  #D#.#C#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
		},
	}

	for _, c := range cases {
		s := parseState(c.start)
		e := parseState(c.expectMove)
		found := false
		for _, m := range s.moves() {
			moved := s.apply(m)
			if moved.equals(e) {
				found = true
				break
			}
		}

		if !found {
			s.Print()
			e.Print()
			t.Errorf("Could not find move ^")
		}
	}
}

func TestLeastEnergy(t *testing.T) {

	type testCases struct {
		name  string
		input string
		cost  int
	}

	cases := []testCases{
		{
			"Example Part 1",
			"" +
				"#############\n" +
				"#...........#\n" +
				"###B#C#B#D###\n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			12521,
		},
		{
			"Real Part 1",
			"" +
				"#############\n" +
				"#...........#\n" +
				"###B#C#C#B###\n" +
				"  #D#D#A#A#  \n" +
				"  #########  \n",
			18051,
		},
		{
			"Example Part 2",
			"" +
				"#############\n" +
				"#...........#\n" +
				"###B#C#B#D###\n" +
				"  #D#C#B#A#  \n" +
				"  #D#B#A#C#  \n" +
				"  #A#D#C#A#  \n" +
				"  #########  \n",
			44169,
		},
		{
			"Real Part 2",
			"" +
				"#############\n" +
				"#...........#\n" +
				"###B#C#C#B###\n" +
				"  #D#C#B#A#  \n" +
				"  #D#B#A#C#  \n" +
				"  #D#D#A#A#  \n" +
				"  #########  \n",
			50245,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := LeastEnergy(c.input)
			if actual != c.cost {
				t.Errorf("expected %depth, got %depth", c.cost, actual)
			}
		})
	}
}
