package generator

import (
	"maze/model"
	"maze/util"
	"maze/model/cell"
)

func BinaryTree (g model.Grid) {
	for c := range g.CellIter() {
		c := c.(*cell.GridCell)
		neighbors := make([]cell.Cell, 0, 2);

		if c.North() != nil {
			neighbors = append(neighbors, c.North());
		}
		if c.East() != nil {
			neighbors = append(neighbors, c.East());
		}

		if len(neighbors) != 0 {
			index := util.RANDOM.Intn(len(neighbors));
			neighbor := neighbors[index];
			c.Link(neighbor, true);
		}
	}
}


