package cell

import "maze/util"

type Cell interface {
	Links() []Cell
	Link(Cell, bool)
	Unlink(Cell, bool)
	IsLinked(Cell) bool

	// Getters & setters -- inherited from baseCell
	Row() int
	Column() int

	Neighbors() []Cell
	Distances() *Distances
}

// Part 1 of an implementation of Djikstra's graph search algorithm as applied to mazes.
func calculateDistances(c Cell) *Distances {
	distances := NewDistances(c);
	frontier := []Cell{c};

	for SliceHasAny(frontier) {
		new_frontier := []Cell{};

		for _, cell := range(frontier) {
			for _, linked := range(cell.Links()) {
				_, ok := distances.Cells[linked]
				if ok {
					continue
				}
				distances.Cells[linked] = distances.Cells[cell] + 1;
				new_frontier = append(new_frontier, linked);
			}
		}
		frontier = new_frontier;
	}
	return distances;
}


/***********************************************************

     Utility functions for working with slices of Cells

************************************************************/

// Check to see if any of the *Cells in this slice are not nil. Akin to Ruby's `any?`
func SliceHasAny (cells []Cell) bool {
	for _, c := range cells  {
		if c != nil {
			return true;
		}
	}
	return false
}

func RandomSliceElement (cells []Cell) Cell {
	return cells[util.RANDOM.Intn(len(cells))];
}

func IndexOf(cells []Cell, cell Cell) int {
	for i, elem := range(cells) {
		if elem == cell {
			return i;
		}
	}
	return -1;
}

func DoesSliceInclude (cells []Cell, cell Cell) bool {
	return IndexOf(cells, cell) != -1;
}
