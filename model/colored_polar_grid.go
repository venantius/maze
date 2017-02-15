package model

import (
	"image/draw"
	"maze/model/cell"
	"image/color"
	"image"
	"maze/sketch"
	"maze/util"
	"fmt"
)

type PolarGrid interface {
	Grid
	drawPolarBackground(*cell.PolarCell, draw.Image, int, int, color.Color)
}

type coloredPolarGrid struct {
	*polarGrid

	distances *cell.Distances
	maximum int
}

func NewColoredPolarGrid(rows int) *coloredPolarGrid {
	return &coloredPolarGrid{
		NewPolarGrid(rows),
		nil,
		0,
	}
}

func (p *coloredPolarGrid) SetDistances(dist *cell.Distances) {
	p.distances = dist;
	_, p.maximum = dist.Max();
}

func (p *coloredPolarGrid) backgroundColorFor(c cell.Cell) color.Color {
	distance, ok := p.distances.Cells[c]
	if !ok {
		return nil
	}
	var intensity float64 = float64(p.maximum-distance) / float64(p.maximum);
	var dark uint8 = uint8(util.Round(255 * intensity));
	var bright uint8 = uint8(128 + util.Round(127*intensity));

	return color.RGBA{dark, bright, bright, 0xff};
}

func (p *coloredPolarGrid) drawPolarBackground(c *cell.PolarCell, img draw.Image, center int, cellSize int, wall color.Color) {
	var theta float64 			= p.calculateTheta(c);
	var inner_radius float64 	= float64(c.Row() * cellSize);
	var theta_ccw float64 		= float64(c.Column()) * theta;
	var theta_cw float64		= float64(c.Column() + 1) * theta;

	p1 := pointFromTheta(center, inner_radius, theta_ccw);
	p2 := pointFromTheta(center, inner_radius, theta_cw);

	points := []image.Point{p1, p2}

	var p3 image.Point;
	var p4 image.Point;

	if len(c.Outward) > 0 {
		outerCell := c.Outward[0].(*cell.PolarCell);

		theta = p.calculateTheta(outerCell);
		inner_radius = float64(outerCell.Row() * cellSize);
		theta_ccw = float64(outerCell.Column()) * theta;
		theta_cw = float64(outerCell.Column() + 1) * theta;

		if len(c.Outward) == 1 {
			// Just a normal outer cell
			p3 = pointFromTheta(center, inner_radius, theta_cw);
			p4 = pointFromTheta(center, inner_radius, theta_ccw);
		} else if len(c.Outward) == 2 {
			// Adaptive division has occurred
			theta_ccw2 := float64(outerCell.Column() + 2) * theta;
			p5 := pointFromTheta(center, inner_radius, theta_ccw2);
			p3 = pointFromTheta(center, inner_radius, theta_cw);
			p4 = pointFromTheta(center, inner_radius, theta_ccw);
			points = append(points, p5);
		} else {
			// Center cell.
			points = []image.Point{};
			fmt.Println(c);
			fmt.Println(c.Outward);
			for _, cel := range(c.Outward) {
				theta = p.calculateTheta(cel.(*cell.PolarCell));
				inner_radius = float64(cel.Row() * cellSize);
				theta_ccw = float64(cel.Column()) * theta;
				points = append(points, pointFromTheta(center, inner_radius, theta_ccw));
			}
			col := p.backgroundColorFor(c);
			sketch.DrawQuadrilateral(points, img, col, wall);
			return;
		}
	} else {
		// Outer-most ring.
		var outer_radius float64 = float64((c.Row() + 1) * cellSize);

		p3 = pointFromTheta(center, outer_radius, theta_cw);
		p4 = pointFromTheta(center, outer_radius, theta_ccw);
	}
	points = append(points, p3);
	points = append(points, p4);
	col := p.backgroundColorFor(c);
	sketch.DrawQuadrilateral(points, img, col, wall);
}

func (p *coloredPolarGrid) ToPNG(filename string, cellSize int) {
	polarGridToPNG(p, filename, cellSize);
}
