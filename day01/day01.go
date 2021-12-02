package day01

import "AdventCode2020/util"

func countDepthIncreasesInWindow(measurements []int, window int) int {
	minMeasurements := window + 1
	if len(measurements) < minMeasurements {
		return 0
	}

	increases := 0
	for i := window; i < len(measurements); i++ {
		if util.Sum(measurements[i-(window-1):i+1]) > util.Sum(measurements[i-window:i]) {
			increases++
		}
	}

	return increases
}
