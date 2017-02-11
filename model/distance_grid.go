package model

import (
	"strconv"
	"maze/model/cell"
)

type IDistanceGrid interface {
	Grid

	SetDistances(*cell.Distances)
}

type distanceGrid struct {
	*baseGrid

	distances *cell.Distances
}

func NewDistanceGrid(rows int, columns int) *distanceGrid {
	return &distanceGrid{
		NewBaseGrid(rows, columns),
		nil,
	}
}

// Set the distances for this grid.
func (d *distanceGrid) SetDistances(dist *cell.Distances) {
	d.distances = dist;
}

func (d *distanceGrid) contentsOf (c cell.Cell) string {
	if d.distances != nil {
		_, ok := d.distances.Cells[c];
		if ok {
			return strconv.FormatInt(int64(d.distances.Cells[c]), 36);
		}
	}
	return d.baseGrid.contentsOf(c);
}

func (d *distanceGrid) String() string {
	return gridString(d);
}
