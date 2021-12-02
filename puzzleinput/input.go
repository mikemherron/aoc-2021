package puzzleinput

import (
	"AdventCode2020/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadIntsFrom(filePath string) []int {
	lines := ReadLinesFrom(filePath)
	ints := make([]int, 0, len(lines))
	for _, line := range lines {
		ints = append(ints, util.TryParseInt(line))
	}

	return ints
}

func ReadLinesFrom(fileName string) []string {
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
