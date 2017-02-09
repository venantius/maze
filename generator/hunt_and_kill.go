package generator

import (
	"maze/model"
)

func HuntAndKill(g model.Grid) {
	current := g.RandomCell();

	for current != nil {
		unvisited_neighbors := []model.Cell{};

		// In ruby: unvisited_neighbors = current.neighbors.select { |n| n.links.empty? }
		for _, neighbor := range(current.Neighbors()) {
			if len(neighbor.Links()) == 0 {
				unvisited_neighbors = append(unvisited_neighbors, neighbor)
			}
		}

		if model.SliceHasAny(unvisited_neighbors) {
			var neighbor model.Cell = model.RandomSliceElement(unvisited_neighbors);
			current.Link(neighbor, true);
			current = neighbor;
		} else {
			current = nil;

			for cell := range(g.CellIter()) {
				var visited_neighbors []model.Cell = []model.Cell{};
				for _, n := range(cell.Neighbors()) {
					if model.SliceHasAny(n.Links()) {
						visited_neighbors = append(visited_neighbors, n);
					}
				}

				if len(cell.Links()) == 0 && model.SliceHasAny(visited_neighbors) {
					current = cell;
					var neighbor model.Cell = model.RandomSliceElement(visited_neighbors);
					current.Link(neighbor, true);
					break;
				}
			}
		}
	}
}
