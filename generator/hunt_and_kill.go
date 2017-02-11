package generator

import (
	"maze/model"
	"maze/model/cell"
)

func HuntAndKill(g model.Grid) {
	current := g.RandomCell();

	for current != nil {
		unvisited_neighbors := []cell.Cell{};

		// In ruby: unvisited_neighbors = current.neighbors.select { |n| n.links.empty? }
		for _, neighbor := range(current.Neighbors()) {
			if len(neighbor.Links()) == 0 {
				unvisited_neighbors = append(unvisited_neighbors, neighbor)
			}
		}

		if cell.SliceHasAny(unvisited_neighbors) {
			var neighbor cell.Cell = cell.RandomSliceElement(unvisited_neighbors);
			current.Link(neighbor, true);
			current = neighbor;
		} else {
			current = nil;

			for c := range(g.CellIter()) {
				var visited_neighbors []cell.Cell = []cell.Cell{};
				for _, n := range(c.Neighbors()) {
					if cell.SliceHasAny(n.Links()) {
						visited_neighbors = append(visited_neighbors, n);
					}
				}

				if len(c.Links()) == 0 && cell.SliceHasAny(visited_neighbors) {
					current = c;
					var neighbor cell.Cell = cell.RandomSliceElement(visited_neighbors);
					current.Link(neighbor, true);
					break;
				}
			}
		}
	}
}
