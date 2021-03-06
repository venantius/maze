package model

import (
	"fmt"
	"image"
	"image/color"
	"maze/util"
	"maze/model/cell"
)

type coloredGrid struct {
	*baseGrid

	distances *cell.Distances
	maximum   int
}

func NewColoredGrid(rows int, columns int) *coloredGrid {
	return &coloredGrid{
		NewBaseGrid(rows, columns),
		nil,
		0,
	}
}

func (cg *coloredGrid) SetDistances(d *cell.Distances) {
	cg.distances = d
	_, cg.maximum = d.Max()
}

func (cg *coloredGrid) backgroundColorFor(c cell.Cell) color.Color {
	distance, ok := cg.distances.Cells[c]
	if !ok {
		return nil
	}
	var intensity float64 = float64(cg.maximum-distance) / float64(cg.maximum)
	var dark uint8 = uint8(util.Round(255 * intensity))
	var bright uint8 = uint8(128 + util.Round(127*intensity))

	return color.RGBA{dark, bright, bright, 0xff}
}

func (cg *coloredGrid) ToPNG(filename string, size int) {
	cg.orderedDistances()
	gridToPNG(cg, filename, size)
}

// TODO: Return an ordered list of cells by distance.
func (cg *coloredGrid) orderedDistances() []cell.Cell {
	var output []cell.Cell = make([]cell.Cell, 0)

	for cell := range cg.CellIter() {
		output = append(output, cell)
	}
	return output
}

func ToGIF(g Grid, filename string, cellSize int) {
	imgWidth := cellSize * g.Columns()
	imgHeight := cellSize * g.Rows()

	var images []*image.Paletted
	var delays []int

	var steps int = g.Rows() * g.Columns()

	// TODO: Figure out how we want to palette this.
	// Possibly we generate all of the images first...and then?
	var pallette = []color.Color{}

	for _, mode := range []int{BACKGROUNDS, WALLS} {
		for c := range g.CellIter() {
			img := image.NewPaletted(image.Rect(0, 0, imgWidth+1, imgHeight+1), pallette)
			fmt.Println(mode, c, img)

		}
	}

	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, imgWidth+1, imgHeight+1), pallette)
		images = append(images, img) // TODO: Maybe move this to the end, after image generation?
		delays = append(delays, 0)
	}

}
