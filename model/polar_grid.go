package model

import (
	"image/color"
	"image"
	"maze/sketch"
	"github.com/llgcode/draw2d/draw2dimg"
	"math"
	"fmt"
	"maze/util"
)

/* Chapter 7 */

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

func (p *polarGrid) prepareGrid() [][]Cell {
	rows := make([][]Cell, p.GetRows());

	rowHeight := 1.0 / float64(p.rows);
	rows[0] = []Cell{NewPolarCell(0, 0)};

	// I'm worried that I'm doing this wrong, but we'll go with it for now.
	// In ruby this is (1..@rows).each do |row|
	for row := 1; row < p.rows; row++ {

		radius := float64(row) / float64(p.rows);
		circumference := 2 * math.Pi * radius;

		previousCount := len(rows[row - 1]);
		estimatedCellWidth := circumference / float64(previousCount);
		ratio := util.Round(estimatedCellWidth / float64(rowHeight));

		cells := previousCount * ratio;
		tmp := make([]Cell, cells);
		for i, _ := range(tmp) {
			tmp[i] = NewPolarCell(row, i);
		}
		rows[row] = tmp;
	}

	return rows;
}

func (p *polarGrid) ToPNG(filename string, cellSize int) {
	var imgSize int = 2 * p.rows * cellSize;

	// background := color.RGBA{0xff, 0xff, 0xff, 0xff};
	wall := color.RGBA{0x44, 0x44, 0x44, 0xff};

	img := image.NewRGBA(image.Rect(0, 0, imgSize + 1, imgSize + 1));
	var center int = imgSize / 2;
	fmt.Println(center);

	// First, set a white background.
	sketch.DrawRectangle(0, 0, imgSize + 1, imgSize + 1, img, color.White);

	for cell := range p.CellIter() {
		var theta float64 			= 2 * math.Pi / float64(len(p.grid[cell.Row()]));
		var inner_radius float64 	= float64(cell.Row() * cellSize);
		var outer_radius float64 	= float64((cell.Row() + 1) * cellSize);
		var theta_ccw float64 		= float64(cell.Col()) * theta;
		var theta_cw float64		= float64(cell.Col() + 1) * theta;

		var ax int = center + int(inner_radius * math.Cos(theta_ccw));
		var ay int = center + int(inner_radius * math.Sin(theta_ccw));
		// var bx int = center + int(outer_radius * math.Cos(theta_ccw));
		// var by int = center + int(outer_radius * math.Sin(theta_ccw));

		var cx int = center + int(inner_radius * math.Cos(theta_cw));
		var cy int = center + int(inner_radius * math.Sin(theta_cw));
		var dx int = center + int(outer_radius * math.Cos(theta_cw));
		var dy int = center + int(outer_radius * math.Sin(theta_cw));

		if (!cell.IsLinked(cell.North())) {
			sketch.DrawLine(ax, ay, cx, cy, img, wall);
		}
		if (!cell.IsLinked(cell.East())) {
			sketch.DrawLine(cx, cy, dx, dy, img, wall)
		}
	}

	draw2dimg.SaveToPngFile(filename, img);
}