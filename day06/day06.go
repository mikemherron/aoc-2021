package day06

type StartState struct {
	state int
	days  int
}

func SimulateLanternFish(starting []int, days int) int64 {

	total := int64(0)
	solutions := make(map[StartState]int64)
	for i := 0; i < len(starting); i++ {
		total += simulateSingleFish(starting[i], 0, days, solutions)
	}

	return total
}

func simulateSingleFish(state int, start int, days int, solutions map[StartState]int64) int64 {
	k := StartState{state, days - start}
	if v, ok := solutions[k]; ok {
		return v
	}

	total := int64(1)
	for i := start; i <= days; i++ {
		if state < 0 {
			total += simulateSingleFish(8, i, days, solutions)
			state = 6
		}
		state--
	}

	solutions[k] = total
	return total
}
