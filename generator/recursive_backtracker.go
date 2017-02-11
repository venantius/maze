package generator

import (
	"maze/model"
	"maze/model/cell"
)

func RecursiveBacktracker(g model.Grid) {
	var start_at cell.Cell = g.RandomCell();

	var stack []cell.Cell = []cell.Cell{start_at};

	for cell.SliceHasAny(stack) {
		current := stack[len(stack)-1]; // grab the last element

		neighbors := []cell.Cell{};
		for _, neighbor := range(current.Neighbors()) {
			if len(neighbor.Links()) == 0 {
				neighbors = append(neighbors, neighbor)
			}
		}

		if len(neighbors) == 0 {
			stack = stack[:len(stack) - 1]
		} else {
			neighbor := cell.RandomSliceElement(neighbors);
			current.Link(neighbor, true);
			stack = append(stack, neighbor);
		}
	}
}
