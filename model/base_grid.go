package model

import (
	"math/rand"
	"image/color"
	"fmt"
)

const NEWLINE_DELIMITER string = "\n";

type baseGrid struct {
	rows 	int
	columns int

	grid 	[][]*cell
}

// Initialize a new grid and return a pointer to it.
func NewBaseGrid(rows int, columns int) *baseGrid {
	var g *baseGrid = &baseGrid{
		rows: rows,
		columns: columns,

		grid: prepareGrid(rows, columns),
	}
	g.configureCells()
	return g
}

// Iterate through the grid and initialize a cell struct for each grid element.
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
func (g *baseGrid) configureCells() {
	for i, row := range(g.grid) {
		for j, _ := range(row) {
			var c *cell = g.grid[i][j];

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

// Returns how many integer columns this grid has
func (g *baseGrid) GetColumns() int {
	return g.columns;
}

// Returns how many integer rows this grid has
func (g *baseGrid) GetRows() int {
	return g.rows;
}

// Retrieve a specific cell within the grid. If the request is for an out-of-bounds cell, returns nil.
// NOTE: This latter capability may only exist to satisfy weird Ruby behavior.
func (g *baseGrid) GetCell(row int, column int) *cell {
	if (row >= 0 && row < g.rows) &&
		(column >= 0 && column < g.columns) {
		return g.grid[row][column]
	}
	return nil
}

// Retrieve a random cell from the grid.
func (g *baseGrid) RandomCell() *cell {
	var row int = rand.Intn(g.rows);
	var column int = rand.Intn(g.columns);
	return g.grid[row][column];
}

// How many cells are in this grid in total?
func (g *baseGrid) Size() int {
	return g.rows * g.columns;
}

func (g *baseGrid) RowIter() <-chan []*cell {
	ch := make(chan []*cell);
	go func () {
		for _, row := range g.grid {
			ch <- row
		}
		close(ch);
	} ();
	return ch;
}

func (g *baseGrid) CellIter() <-chan *cell {
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

/*
Satisfying the Grid interface.
 */

func (g *baseGrid) contentsOf(c *cell) string {
	return " ";
}

func (g *baseGrid) backgroundColorFor(*cell) color.Color {
	fmt.Println("unfortunately...")
	return nil;
}

func (g *baseGrid) String() string {
	return gridString(g);
}

func (g *baseGrid) ToPNG(filename string, size int) {
	gridToPNG(g, filename, size);
}
