package model

import (
	"image/color"
	"image"
	"maze/sketch"
	"github.com/llgcode/draw2d/draw2dimg"
	"math"
	"maze/util"
	"maze/model/cell"
)

/* Chapter 7

A @note about polar grids -- the earlier chapters of the book have algorithms that rely on North(), South(), East() and West()
methods -- these methods are meaningless in the context of a polar grid. Consequently, only algorithms like
Wilson's, Recursive Backtracker, and Aldous-Broder will work with the library in its current state.

*/

type polarGrid struct {
	*baseGrid
}

func NewPolarGrid(rows int) *polarGrid {
	p := &polarGrid{
		&baseGrid{
			rows: rows,
			columns: 1,
		},
	}
	p.grid = p.prepareGrid();
	p.configureCells();
	return p;
}

func (p *polarGrid) prepareGrid() [][]cell.Cell {
	rows := make([][]cell.Cell, p.GetRows());

	rowHeight := 1.0 / float64(p.rows);
	rows[0] = []cell.Cell{cell.NewPolarCell(0, 0)};

	// I'm worried that I'm doing this wrong, but we'll go with it for now.
	// In ruby this is (1...@rows).each do |row|
	for row := 1; row < p.rows; row++ {

		radius := float64(row) / float64(p.rows);
		circumference := 2 * math.Pi * radius;

		previousCount := len(rows[row - 1]);
		estimatedCellWidth := circumference / float64(previousCount);
		ratio := util.Round(estimatedCellWidth / float64(rowHeight));

		cells := previousCount * ratio;
		tmp := make([]cell.Cell, cells);
		for i, _ := range(tmp) {
			tmp[i] = cell.NewPolarCell(row, i);
		}
		rows[row] = tmp;
	}

	return rows;
}

func (p *polarGrid) configureCells() {
	for c := range p.CellIter() {
		var c *cell.PolarCell = c.(*cell.PolarCell)
		row, col := c.Row(), c.Column()
		if row > 0 {
			c.CCW = p.GetCell(row, col + 1);
			c.CW = p.GetCell(row, col - 1);

			var ratio float64 = float64(len(p.grid[row])) / float64(len(p.grid[row-1]));
			var parent *cell.PolarCell = p.grid[row - 1][int(float64(col) / ratio)].(*cell.PolarCell);
			parent.Outward = append(parent.Outward, c);
			c.Inward = parent;
		}
	}
}

func (p *polarGrid) Size() int {
	count := 0
	for range(p.CellIter()) {
		count++
	}
	return count;
}

func (p *polarGrid) GetCell(row int, column int) cell.Cell {
	if row >= 0 && row <= p.rows - 1 {
		c := column % len(p.grid[row]);
		if c < 0 {
			c = c + len(p.grid[row]);
		}
		return p.grid[row][c];
	}
	return nil;
}

func (p *polarGrid) RandomCell() cell.Cell {
	var row int = util.RANDOM.Intn(p.rows);
	var col int = util.RANDOM.Intn(len(p.grid[row]));
	return p.grid[row][col];
}

func (p *polarGrid) ToPNG(filename string, cellSize int) {
	var imgSize int = 2 * p.rows * cellSize;

	// background := color.RGBA{0xff, 0xff, 0xff, 0xff};
	wall := color.RGBA{0x44, 0x44, 0x44, 0xff};

	img := image.NewRGBA(image.Rect(0, 0, imgSize + 1, imgSize + 1));
	var center int = imgSize / 2;

	// First, set a white background.
	sketch.DrawRectangle(0, 0, imgSize + 1, imgSize + 1, img, color.White);

	for c := range p.CellIter() {
		// Skip the inner-most cell.
		c := c.(*cell.PolarCell);
		if c.Row() == 0 {
			continue;
		}
		var theta float64 			= 2 * math.Pi / float64(len(p.grid[c.Row()]));
		var inner_radius float64 	= float64(c.Row() * cellSize);
		var outer_radius float64 	= float64((c.Row() + 1) * cellSize);
		var theta_ccw float64 		= float64(c.Column()) * theta;
		var theta_cw float64		= float64(c.Column() + 1) * theta;

		var ax int = center + int(inner_radius * math.Cos(theta_ccw));
		var ay int = center + int(inner_radius * math.Sin(theta_ccw));
		var bx int = center + int(outer_radius * math.Cos(theta_ccw));
		var by int = center + int(outer_radius * math.Sin(theta_ccw));

		var cx int = center + int(inner_radius * math.Cos(theta_cw));
		var cy int = center + int(inner_radius * math.Sin(theta_cw));
		var dx int = center + int(outer_radius * math.Cos(theta_cw));
		var dy int = center + int(outer_radius * math.Sin(theta_cw));

		if (!c.IsLinked(c.Inward)) {
			sketch.DrawLine(ax, ay, cx, cy, img, wall);
		}
		if (!c.IsLinked(c.CCW)) {
			sketch.DrawLine(cx, cy, dx, dy, img, wall)
		}

		if c.Row() == p.rows - 1 {
			sketch.DrawLine(bx, by, dx, dy, img, wall);
		}
	}

	draw2dimg.SaveToPngFile(filename, img);
}