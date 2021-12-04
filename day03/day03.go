package day03

import (
	"AdventCode2020/util"
)

type diagnostics struct {
	bits   int
	values []int
}

func (d *diagnostics) isMostBitsSet(pos int) bool {
	c := 0
	for _, v := range d.values {
		if v&(1<<pos) == 0 {
			c--
		} else {
			c++
		}
	}

	return c >= 0
}

func (d *diagnostics) reduce(pos int, set bool) {
	if len(d.values) == 1 {
		return
	}

	tmp := d.values[:0]
	for _, v := range d.values {
		if v&(1<<pos) > 0 == set {
			tmp = append(tmp, v)
		}
	}

	d.values = tmp
}

func (d *diagnostics) first() int {
	return d.values[0]
}

func (d *diagnostics) copy() *diagnostics {
	valuesCopy := make([]int, len(d.values))
	copy(valuesCopy, d.values)

	return &diagnostics{d.bits, valuesCopy}
}

func newDiagnostics(readings []string) *diagnostics {
	v := make([]int, len(readings))
	for i, d := range readings {
		v[i] = util.TryParseBinary(d)
	}

	return &diagnostics{len(readings[0]), v}
}

func ProcessDiagnostics(readings []string) (int, int) {
	gamma, epsilon := 0, 0

	all := newDiagnostics(readings)
	o2 := all.copy()
	co2 := all.copy()

	for i := all.bits - 1; i >= 0; i-- {
		if all.isMostBitsSet(i) {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}

		o2.reduce(i, o2.isMostBitsSet(i))
		co2.reduce(i, !co2.isMostBitsSet(i))
	}

	return gamma * epsilon, co2.first() * o2.first()
}
