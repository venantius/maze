package generator

import (
	"maze/model"
)

func RecursiveBacktracker(g model.Grid) {
	var start_at model.Cell = g.RandomCell();

	var stack []model.Cell = []model.Cell{start_at};

	for model.SliceHasAny(stack) {
		current := stack[len(stack)-1]; // grab the last element

		neighbors := []model.Cell{};
		for _, neighbor := range(current.Neighbors()) {
			if len(neighbor.Links()) == 0 {
				neighbors = append(neighbors, neighbor)
			}
		}

		if len(neighbors) == 0 {
			stack = stack[:len(stack) - 1]
		} else {
			neighbor := model.RandomSliceElement(neighbors);
			current.Link(neighbor, true);
			stack = append(stack, neighbor);
		}
	}
}
