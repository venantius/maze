package generator

import (
	"maze/model"
	"maze/util"
	"fmt"
)

func Wilsons(g model.Grid) {
	unvisited := []*model.Cell{};

	// Mark all cells as unvisited.
	for cell := range(g.CellIter()) {
		unvisited = append(unvisited, cell);
	}

	fmt.Println(len(unvisited));

	// Pick a random unvisited node and mark it visited.
	index := util.RANDOM.Intn(len(unvisited));
	unvisited = append(unvisited[:index], unvisited[index+1:]...)
	fmt.Println(len(unvisited));


	// While any cells are unvisited...
	for model.SliceHasAny(unvisited) {

		// Pick a random starting cell.
		cell := unvisited[util.RANDOM.Intn(len(unvisited))];
		path := []*model.Cell{cell};

		// Keep building a path until the cell isn't an unvisited one.
		for model.DoesSliceInclude(unvisited, cell) {
			cell = model.RandomSliceElement(cell.Neighbors());

			position := model.IndexOf(path, cell);
			if position != -1 {
				path = path[0:position]
			} else {
				path = append(path, cell);
			}

		}

		fmt.Println(len(unvisited));


		for i := 0; i < len(path) - 1; i++ {
			path[i].Link(path[i+1], true);
			index = model.IndexOf(unvisited, path[i]);
			unvisited = append(unvisited[:index], unvisited[index+1:]...)
		}
	}
}
