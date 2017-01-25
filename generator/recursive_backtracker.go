package generator

import "maze/model"

func RecursiveBacktracker(g model.Grid) {
	var start_at *model.Cell = g.RandomCell();

	var stack []*model.Cell = []*model.Cell{start_at};

	for model.SliceHasAny(stack) {
	}
}
