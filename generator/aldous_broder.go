package generator

import (
	"maze/model"
	"maze/util"
)

func AldousBroder (g model.Grid) {
	var cell *model.Cell = g.RandomCell();

	var unvisited int = g.Size() - 1;

	for unvisited > 0 {
		var neighbor *model.Cell = cell.Neighbors()[util.RANDOM.Intn(len(cell.Neighbors()))];

		if len(neighbor.Links()) == 0 {
			cell.Link(neighbor, true);
			unvisited -= 1
		}

		cell = neighbor;
	}
}