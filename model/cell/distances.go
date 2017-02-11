package cell

type Distances struct {
	root  Cell
	Cells map[Cell]int
}

func NewDistances(root Cell) *Distances {
	c := make(map[Cell]int);
	c[root] = 0;
	return &Distances {
		root: root,
		Cells: c,
	}
}

func (d *Distances) CellIter() <-chan Cell {
	ch := make(chan Cell, 1);
	go func() {
		for c, _ := range(d.Cells) {
			ch <- c
		}
		close(ch);
	} ();
	return ch;
}

// Return a new Distances with the shortest path between the root *cell of this distances and the goal *cell, provided.
func (d *Distances) PathTo(goal Cell) *Distances {
	current := goal;

	breadcrumbs := NewDistances(current);
	breadcrumbs.Cells[current] = d.Cells[current];

	for current != d.root {
		for _, neighbor := range(current.Links()) {
			if d.Cells[neighbor] < d.Cells[current] {
				breadcrumbs.Cells[neighbor] = d.Cells[neighbor];
				current = neighbor;
			}
		}
	}

	return breadcrumbs
}

// Which cell is furthest away from the root?
func (d *Distances) Max() (Cell, int) {
	maxDistance := 0;
	maxCell := d.root;

	for cell, distance := range(d.Cells) {
		if distance > maxDistance {
			maxDistance = distance;
			maxCell = cell;
		}
	}
	return maxCell, maxDistance;
}
