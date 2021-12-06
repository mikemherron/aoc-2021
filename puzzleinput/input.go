package puzzleinput

import (
	"AdventCode2021/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadCommaSeperatedInts(filePath string) []int {
	lines := ReadLines(filePath)
	return util.SplitByCommaToInt(lines[0])
}

func ReadIntLines(filePath string) []int {
	lines := ReadLines(filePath)
	ints := make([]int, 0, len(lines))
	for _, line := range lines {
		ints = append(ints, util.TryParseInt(line))
	}

	return ints
}

func ReadLines(fileName string) []string {
	pwd, e := os.Getwd()
	if e != nil {
		panic(e)
	}

	all, e := ioutil.ReadFile(filepath.Join(pwd, fileName))
	if e != nil {
		panic(e)
	}

	return strings.Split(string(all), "\n")
}
