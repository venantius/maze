package model

import (
	"image"
	"bytes"
	"fmt"
	"image/color"
	"maze/sketch"
	"image/draw"
	"github.com/llgcode/draw2d/draw2dimg"
)

const (
	BACKGROUNDS = iota
	WALLS
)

type Grid interface{
	fmt.Stringer

	// Getters
	GetColumns() int
	GetRows() int

	RandomCell() *Cell
	Size() int

	// Iterators
	RowIter() <-chan []*Cell
	CellIter() <-chan *Cell

	// Misc
	Deadends() []*Cell

	// Printing the maze
	contentsOf(*Cell) string
	backgroundColorFor(*Cell) color.Color
	ToPNG(filename string, size int)
}


// ASCII representation
func gridString(g Grid) string {
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
			if cell.IsLinked(cell.East) {
				east_boundary = " "
			} else {
				east_boundary = "|"
			}

			top += body + east_boundary

			// three spaces below, too >> >...<
			var south_boundary string
			if cell.IsLinked(cell.South) {
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
func gridToPNG(g Grid, filename string, cellSize int) {
	imgWidth := cellSize * g.GetColumns();
	imgHeight := cellSize * g.GetRows();

	// background := color.RGBA{0xff, 0xff, 0xff, 0xff};
	// wall := color.RGBA{0x44, 0x44, 0x44, 0xff};

	img := image.NewRGBA(image.Rect(0, 0, imgWidth + 1, imgHeight + 1));

	// First, set a white background.
	sketch.DrawRectangle(0, 0, imgWidth + 1, imgHeight + 1, img, color.White);

	for _, mode := range []int{BACKGROUNDS, WALLS} {
		for c := range g.CellIter() {
			drawGrid(mode, g, c, img, cellSize);
		}
	}

	draw2dimg.SaveToPngFile(filename, img);
}

func drawGrid(drawMode int, g Grid, c *Cell, img draw.Image, cellSize int) {
	x1 := c.column * cellSize;
	y1 := c.row * cellSize;
	x2 := (c.column + 1) * cellSize;
	y2 := (c.row + 1) * cellSize;

	if drawMode == BACKGROUNDS {
		col := g.backgroundColorFor(c);
		if col != nil {
			sketch.DrawRectangle(x1, y1, x2, y2, img, col);
		}
	} else {
		if c.North == nil {
			sketch.DrawLine(x1, y1, x2, y1, img, color.Black);
		}
		if c.West == nil {
			sketch.DrawLine(x1, y1, x1, y2, img, color.Black);
		}
		if !c.IsLinked(c.East) {
			sketch.DrawLine(x2, y1, x2, y2, img, color.Black);
		}
		if !c.IsLinked(c.South) {
			sketch.DrawLine(x1, y2, x2, y2, img, color.Black);
		}
	}
}
