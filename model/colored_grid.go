package model

import (
	"image/color"
	"image"
	"maze/util"
)

type coloredGrid struct {
	*baseGrid

	distances *Distances
	maximum int
}

func NewColoredGrid(rows int, columns int) *coloredGrid {
	return &coloredGrid{
		NewBaseGrid(rows, columns),
		nil,
		0,
	}
}

func (cg *coloredGrid) SetDistances(d *Distances) {
	cg.distances = d;
	_, cg.maximum = d.Max();
}

func (cg *coloredGrid) backgroundColorFor(c *Cell) color.Color {
	distance, ok := cg.distances.cells[c];
	if !ok {
		return nil;
	}
	var intensity float64 = float64(cg.maximum - distance) / float64(cg.maximum);
	var dark uint8 = uint8(util.Round(255 * intensity));
	var bright uint8 = uint8(128 + util.Round(127 * intensity));

	return color.RGBA{dark, bright, dark, 0xff};
}

func (cg *coloredGrid) ToPNG(filename string, size int) {
	gridToPNG(cg, filename, size);
}

func ToGIF(g Grid, filename string, cellSize int) {
	imgWidth := cellSize * g.GetColumns();
	imgHeight := cellSize * g.GetRows();

	var images []*image.Paletted;
	var delays []int;

	var steps int = g.GetRows() * g.GetColumns();

	// TODO: Figure out how we want to palette this.
	// Possibly we generate all of the images first...and then?
	var pallette = []color.Color{};

	for step := 0; step < steps; step++ {
		img := image.NewPaletted(image.Rect(0, 0, imgWidth + 1, imgHeight + 1), pallette);
		images = append(images, img); // TODO: Maybe move this to the end, after image generation?
		delays = append(delays, 0);
	}

}


