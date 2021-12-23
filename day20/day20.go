package day20

import (
	"AdventCode2021/util"
	"AdventCode2021/util/grid"
	"strconv"
)

const (
	PixelOn  = "#"
	PixelOff = "."
)

func EnhancedPixels(input []string, passes int) int {

	enhancement, image := input[0], newImage(input[2:])

	infiniteEven := string(enhancement[0])
	infiniteOdd := PixelOff

	for i := 0; i < passes; i++ {
		image = enhance(image, enhancement)
		if i%2 == 0 {
			image.infinite = infiniteEven
		} else {
			image.infinite = infiniteOdd
		}
	}

	return image.count()
}

func enhance(input *image, enhancement string) *image {
	input = input.grow()
	output := input.copy()

	input.pixels.VisitAll(func(p grid.Pos, v *int) {
		var decodeBinary string
		input.pixels.VisitAdjacentWithFiller(p.Row, p.Col, valFromPixel(input.infinite), grid.Surrounding,
			func(p grid.Pos, v *int) {
				decodeBinary += strconv.Itoa(*v)
			},
		)

		output.setPixel(p.Row, p.Col, valFromPixel(string(enhancement[util.TryParseBinary(decodeBinary)])))
	})

	return output
}

type image struct {
	pixels   *grid.Grid
	infinite string
}

func (i *image) grow() *image {
	imageCopy := image{
		pixels:   i.pixels.Grow(valFromPixel(i.infinite)),
		infinite: i.infinite,
	}

	return &imageCopy
}

func (i *image) copy() *image {
	imageCopy := image{
		pixels:   i.pixels.Copy(),
		infinite: i.infinite,
	}

	return &imageCopy
}

func (i *image) setPixel(r int, c int, v int) {
	g := i.pixels
	(*g)[r][c] = v
}

func (i *image) count() int {
	count := 0
	i.pixels.VisitAll(func(p grid.Pos, v *int) {
		count += *v
	})
	return count
}

func newImage(lines []string) *image {
	g := grid.NewGrid(lines, valFromPixel)
	i := image{pixels: &g, infinite: PixelOff}

	return &i
}

func valFromPixel(p string) int {
	if p == PixelOn {
		return 1
	}
	return 0
}
