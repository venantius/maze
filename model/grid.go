package model

import (
	"image"
	"bytes"
	"fmt"
	"image/color"
	"maze/sketch"
	"image/draw"
	"github.com/llgcode/draw2d/draw2dimg"
	"maze/model/cell"
)

const (
	BACKGROUNDS = iota
	WALLS
)

type Grid interface{
	fmt.Stringer

	// Initializer helpers
	prepareGrid() [][]cell.Cell

	// Getters
	Columns() int
	Rows() int
	Grid() [][]cell.Cell

	RandomCell() cell.Cell
	Size() int

	// Iterators
	RowIter() <-chan []cell.Cell
	CellIter() <-chan cell.Cell

	// Misc
	Deadends() []cell.Cell

	// Printing the maze
	contentsOf(cell.Cell) string
	backgroundColorFor(cell.Cell) color.Color
	ToPNG(filename string, size int)
}

// Iterate through the grid and initialize a cell struct for each grid element.
func prepareGrid(g Grid) [][]cell.Cell {
	grid := make([][]cell.Cell, g.Rows())
	for i, _ := range(grid) {
		grid[i] = make([]cell.Cell, g.Columns());
		for j, _ := range(grid[i]) {
			grid[i][j] = cell.NewGridCell(i, j);
		}
	}
	return grid;
}

// ASCII representation
func gridString(g Grid) string {
	var output bytes.Buffer

	output.WriteString("+");
	for i := 0; i < g.Columns(); i++ {
		output.WriteString("---+");

	}
	output.WriteString(NEWLINE_DELIMITER);

	for row := range g.RowIter() {
		top := "|"
		bottom := "+"

		for _, c := range row {
			c := c.(*cell.GridCell); // Lazy type cast.
			if c == nil {
				c = cell.NewGridCell(-1, -1) // TODO: I think this is un-necessary
			}

			body := fmt.Sprintf(" %v ", g.contentsOf(c));
			var east_boundary string
			if c.IsLinked(c.East()) {
				east_boundary = " "
			} else {
				east_boundary = "|"
			}

			top += body + east_boundary

			// three spaces below, too >> >...<
			var south_boundary string
			if c.IsLinked(c.South()) {
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
	imgWidth := cellSize * g.Columns();
	imgHeight := cellSize * g.Rows();

	// background := color.RGBA{0xff, 0xff, 0xff, 0xff};
	// wall := color.RGBA{0x44, 0x44, 0x44, 0xff};

	img := image.NewRGBA(image.Rect(0, 0, imgWidth + 1, imgHeight + 1));

	// First, set a white background.
	sketch.DrawRectangle(0, 0, imgWidth + 1, imgHeight + 1, img, color.White);

	for _, mode := range []int{BACKGROUNDS, WALLS} {
		for c := range g.CellIter() {
			fmt.Println(c);
			c := c.(*cell.GridCell); // lazy type cast
			drawGrid(mode, g, c, img, cellSize);
		}
	}

	draw2dimg.SaveToPngFile(filename, img);
}

func drawGrid(drawMode int, g Grid, c *cell.GridCell, img draw.Image, cellSize int) {
	x1 := c.Column() * cellSize;
	y1 := c.Row() * cellSize;
	x2 := (c.Column() + 1) * cellSize;
	y2 := (c.Row() + 1) * cellSize;

	if drawMode == BACKGROUNDS {
		col := g.backgroundColorFor(c);
		if col != nil {
			sketch.DrawRectangle(x1, y1, x2, y2, img, col);
		}
	} else {
		if c.North() == nil {
			sketch.DrawLine(x1, y1, x2, y1, img, color.Black);
		}
		if c.West() == nil {
			sketch.DrawLine(x1, y1, x1, y2, img, color.Black);
		}
		if !c.IsLinked(c.East()) {
			sketch.DrawLine(x2, y1, x2, y2, img, color.Black);
		}
		if !c.IsLinked(c.South()) {
			sketch.DrawLine(x1, y2, x2, y2, img, color.Black);
		}
	}
}
