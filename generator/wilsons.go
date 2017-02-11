package generator

import (
	"maze/model"
	"maze/util"
	"maze/model/cell"
)

func Wilsons(g model.Grid) {
	unvisited := []cell.Cell{};

	// Mark all cells as unvisited.
	for c := range(g.CellIter()) {
		unvisited = append(unvisited, c);
	}

	// Pick a random unvisited node and mark it visited.
	index := util.RANDOM.Intn(len(unvisited));
	unvisited = append(unvisited[:index], unvisited[index+1:]...)


	// While any cells are unvisited...
	for cell.SliceHasAny(unvisited) {

		// Pick a random starting cell.
		c := unvisited[util.RANDOM.Intn(len(unvisited))];
		path := []cell.Cell{c};

		// Keep building a path until the cell isn't an unvisited one.
		for cell.DoesSliceInclude(unvisited, c) {
			c = cell.RandomSliceElement(c.Neighbors());

			position := cell.IndexOf(path, c);
			if position != -1 {
				path = path[0:position+1]
			} else {
				path = append(path, c);
			}

		}

		for i := 0; i < len(path) - 1; i++ {
			path[i].Link(path[i+1], true);
			index = cell.IndexOf(unvisited, path[i]);
			unvisited = append(unvisited[:index], unvisited[index+1:]...)
		}
	}
}
