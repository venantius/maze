package model

import (
	"image"
	"bytes"
	"fmt"
	"image/color"
	"github.com/llgcode/draw2d/draw2dimg"
	"mazes/sketch"
	"image/draw"
)

type grid interface{
	fmt.Stringer

	// Getters
	GetColumns() int
	GetRows() int

	// Iterators
	RowIter() <-chan []*cell
	CellIter() <-chan *cell

	// Printing the maze
	contentsOf(*cell) string
	backgroundColorFor(*cell) color.Color
	ToPNG(filename string, size int)
}


// ASCII representation
func gridString(g grid) string {
	var output bytes.Buffer

	output.WriteString("+");
	for i := 0; i < g.GetColumns(); i++ {
		output.WriteString("---+");

	}
	output.WriteString(NEWLINE_DELIMITER);

	for row := range g.RowIter() {
		top := "|"
		bottom := "+"

		for _, cell := range row {
			if cell == nil {
				cell = NewCell(-1, -1) // TODO: I think this is un-necessary
			}

			body := fmt.Sprintf(" %v ", g.contentsOf(cell));
			var east_boundary string
			if cell.IsLinked(cell.east) {
				east_boundary = " "
			} else {
				east_boundary = "|"
			}

			top += body + east_boundary

			// three spaces below, too >> >...<
			var south_boundary string
			if cell.IsLinked(cell.south) {
				south_boundary = "   "
			} else {
				south_boundary = "---"
			}

			bottom += south_boundary + "+"
		}
		output.WriteString(top + NEWLINE_DELIMITER);
		output.WriteString(bottom + NEWLINE_DELIMITER);
	}
	return output.String();
}

// PNG representation
func gridToPNG(g grid, filename string, cellSize int) {
	imgWidth := cellSize * g.GetColumns();
	imgHeight := cellSize * g.GetRows();

	// background := color.RGBA{0xff, 0xff, 0xff, 0xff};
	// wall := color.RGBA{0x44, 0x44, 0x44, 0xff};

	img := image.NewRGBA(image.Rect(0, 0, imgWidth + 1, imgHeight + 1));

	// First, set the background
	for x := 0; x <= imgWidth; x++ {
		for y := 0; y <= imgHeight; y++ {
			img.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 0xff})
		}
	}

	BACKGROUNDS := "backgrounds";
	WALLS := "WALLS"

	for _, mode := range []string{BACKGROUNDS, WALLS} {
		for c := range g.CellIter() {
			drawGrid(mode, g, c, img, cellSize);
		}
	}

	draw2dimg.SaveToPngFile(filename, img);
}

var BACKGROUNDS string = "backgrounds"
var WALLS string = "WALLS"

// TODO: Turn MODE into an enum
func drawGrid(mode string, g grid, c *cell, img draw.Image, cellSize int) {
	x1 := c.column * cellSize;
	y1 := c.row * cellSize;
	x2 := (c.column + 1) * cellSize;
	y2 := (c.row + 1) * cellSize;

	if mode == BACKGROUNDS {
		fmt.Println("Here!")
		color := g.backgroundColorFor(c);
		fmt.Println(color);
		if color != nil {
			sketch.DrawRectangle(x1, y1, x2, y2, img, color);
		}
	} else {
		if c.north == nil {
			sketch.DrawLine(x1, y1, x2, y1, img, color.Black);
		}
		if c.west == nil {
			sketch.DrawLine(x1, y1, x1, y2, img, color.Black);
		}
		if !c.IsLinked(c.east) {
			sketch.DrawLine(x2, y1, x2, y2, img, color.Black);
		}
		if !c.IsLinked(c.south) {
			sketch.DrawLine(x1, y2, x2, y2, img, color.Black);
		}
	}
}
