package model

import (
	"strconv"
)

type distanceGrid struct {
	*baseGrid

	distances *Distances
}

func NewDistanceGrid(rows int, columns int) *distanceGrid {
	return &distanceGrid{
		NewBaseGrid(rows, columns),
		nil,
	}
}

// Set the distances for this grid.
func (d *distanceGrid) SetDistances(dist *Distances) {
	d.distances = dist;
}

func (d *distanceGrid) contentsOf (c *cell) string {
	if d.distances != nil {
		_, ok := d.distances.cells[c];
		if ok {
			return strconv.FormatInt(int64(d.distances.cells[c]), 36);
		}
	}
	return d.baseGrid.contentsOf(c);
}

func (d *distanceGrid) String() string {
	return GridString(d);
}
