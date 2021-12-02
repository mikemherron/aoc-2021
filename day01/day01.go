package day01

import "AdventCode2020/util"

func countDepthIncreasesInWindow(measurements []int, w int) int {
	minMeasurements := w + 1
	if len(measurements) < minMeasurements {
		return 0
	}

	increases := 0
	for i := w; i < len(measurements); i++ {
		if util.Sum(measurements[i-(w-1):i+1]) > util.Sum(measurements[i-w:i]) {
			increases++
		}
	}

	return increases
}
