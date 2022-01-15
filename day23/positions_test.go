package day23

import "testing"

func TestMakePositions(t *testing.T) {

	p := makePositions(4, 0, 4, 1)
	p.assert()
	positions := p.get()
	if len(positions) != 2 {
		t.Errorf("Should be 2 length")
	}

	if positions[0] != [2]uint{4, 0} || positions[1] != [2]uint{4, 1} {
		t.Errorf("Expected {10,0},{2,4}, got %v", positions)
	}

	assert(t, p.occupied(0, 0) == false, "not set")
	assert(t, p.occupied(1, 0) == false, "not set")
	assert(t, p.occupied(4, 0), "should be  set")
	assert(t, p.occupied(4, 1), "should be  set")
	assert(t, p.occupied(10, 0) == false, "not set")

	p.move(4, 0, 6, 2)
	p.assert()
	assert(t, p.occupied(4, 0) == false, "not set")
	assert(t, p.occupied(6, 2), "should be  set")
	assert(t, p.occupied(4, 1), "should be  set")

	p.move(6, 2, 10, 0)
	p.assert()
	assert(t, p.occupied(6, 2) == false, "not set")
	assert(t, p.occupied(10, 0), "should be  set")

	positions = p.get()
	if positions[0] != [2]uint{10, 0} || positions[1] != [2]uint{4, 1} {
		t.Errorf("Expected {10,0},{4,1}, got %v", positions)
	}

	p.add(4, 4)
	positions = p.get()
	if positions[2] != [2]uint{4, 4} {
		t.Errorf("Expected {4,4} got %v", positions[2])
	}
}

func assert(t *testing.T, b bool, s string) {
	if !b {
		t.Errorf(s)
	}
}
