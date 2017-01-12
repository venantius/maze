package model

import (
	_ "fmt"
	"image"
	_ "image/color"
	"math/rand"
	"bytes"
	"fmt"
	"image/color"
	"github.com/llgcode/draw2d/draw2dimg"
	"mazes/sketch"
)

const NEWLINE_DELIM string = "\n";

type grid struct {
	rows 	int
	columns int

	grid 	[][]*cell
}

// Initialize a new grid and return a pointer to it.
func NewGrid(rows int, columns int) *grid {
	var g *grid = &grid {
		rows: rows,
		columns: columns,

		grid: prepareGrid(rows, columns),
	}
	g.configureCells()
	return g
}

// Iterate through the
func prepareGrid(rows int, columns int) [][]*cell {
	grid := make([][]*cell, rows)
	for i, _ := range(grid) {
		column := make([]*cell, columns)
		grid[i] = column
		for j, _ := range(column) {
			c := NewCell(i, j);
			grid[i][j] = c
		}
	}
	return grid
}

// This iterates through each cell in the grid, and for each cell attempts to set a north, east, south and west cell.
// If the cell is at one of the grid's edges, it does not set the neighboring cell (leaving a nil pointer in place).
func (g *grid) configureCells() {
	for i, row := range(g.grid) {
		for j, _ := range(row) {
			var c *cell = g.grid[i][j];

			if c.row + 1 < g.rows {

			}

			if c.row - 1 >= 0 {
				c.north = g.grid[c.row - 1][c.column]
			}
			if c.column + 1 < g.columns {
				c.east  = g.grid[c.row][c.column + 1]
			}
			if c.row + 1 < g.rows {
				c.south = g.grid[c.row + 1][c.column]
			}
			if c.column - 1 >=0 {
				c.west  = g.grid[c.row][c.column - 1]
			}
		}
	}

}

// Retrieve a specific cell within the grid. If the request is for an out-of-bounds cell, returns nil.
// NOTE: This latter capability may only exist to satisfy weird Ruby behavior.
func (g *grid) GetCell(row int, column int) *cell {
	if (row >= 0 && row < g.rows) &&
		(column >= 0 && column < g.columns) {
		return g.grid[row][column]
	}
	return nil
}

// Retrieve a random cell from the grid.
func (g *grid) RandomCell() *cell {
	var row int = rand.Intn(g.rows);
	var column int = rand.Intn(g.columns);
	return g.grid[row][column];
}

// How many cells are in this grid in total?
func (g *grid) Size() int {
	return g.rows * g.columns;
}

func (g *grid) Rows() <-chan []*cell {
	ch := make(chan []*cell);
	go func () {
		for _, row := range g.grid {
			ch <- row
		}
		close(ch);
	} ();
	return ch;
}

func (g *grid) Cells() <-chan *cell {
	ch := make(chan *cell, 1);
	go func () {
		for _, row := range g.grid {
			for _, cell := range row {
				ch <- cell
			}
		}
		close(ch);
	} ();
	return ch;
}

// ASCII representation
func (g *grid) String() string {
	var output bytes.Buffer

	output.WriteString("+");
	for i := 0; i < g.columns; i++ {
		output.WriteString("---+");

	}
	output.WriteString(NEWLINE_DELIM);

	for row := range g.Rows() {
		top := "|"
		bottom := "+"

		for _, cell := range row {
			if cell == nil {
				cell = NewCell(-1, -1) // TODO: I think this is un-necessary
			}

			body := "   " // three spaces
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
			corner := "+"

			bottom += south_boundary + corner
		}
		output.WriteString(top + NEWLINE_DELIM);
		output.WriteString(bottom + NEWLINE_DELIM);
	}
	return output.String();
}

// func setImgBackground

// PNG representation
func (g *grid) ToPNG(cellSize int) {
	imgWidth := cellSize * g.columns;
	imgHeight := cellSize * g.rows;

	// background := color.RGBA{0xff, 0xff, 0xff, 0xff};
	// wall := color.RGBA{0x44, 0x44, 0x44, 0xff};

	img := image.NewRGBA(image.Rect(0, 0, imgWidth + 1, imgHeight + 1));

	// First, set the background
	for x := 0; x <= imgWidth; x++ {
		for y := 0; y <= imgHeight; y++ {
			img.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 0xff})
		}
	}

	for c := range g.Cells() {
		x1 := c.column * cellSize;
		y1 := c.row * cellSize;
		x2 := (c.column + 1) * cellSize;
		y2 := (c.row + 1) * cellSize;

		fmt.Println(float64(x1), float64(y1), float64(x2), float64(y2));
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

	draw2dimg.SaveToPngFile("derp.png", img);
}
