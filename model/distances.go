package model

type Distances struct {
	root	*Cell
	cells	map[*Cell]int
}

func NewDistances(root *Cell) *Distances {
	c := make(map[*Cell]int);
	c[root] = 0;
	return &Distances {
		root: root,
		cells: c,
	}
}

func (d *Distances) Cells () <-chan *Cell {
	ch := make(chan *Cell, 1);
	go func() {
		for c, _ := range(d.cells) {
			ch <- c
		}
		close(ch);
	} ();
	return ch;
}

// Return a new Distances with the shortest path between the root *cell of this distances and the goal *cell, provided.
func (d *Distances) PathTo(goal *Cell) *Distances {
	current := goal;

	breadcrumbs := NewDistances(current);
	breadcrumbs.cells[current] = d.cells[current];

	for current != d.root {
		for _, neighbor := range(current.Links()) {
			if d.cells[neighbor] < d.cells[current] {
				breadcrumbs.cells[neighbor] = d.cells[neighbor];
				current = neighbor;
			}
		}
	}

	return breadcrumbs
}

// Which cell is furthest away from the root?
func (d *Distances) Max() (*Cell, int) {
	maxDistance := 0;
	maxCell := d.root;

	for cell, distance := range(d.cells) {
		if distance > maxDistance {
			maxDistance = distance;
			maxCell = cell;
		}
	}
	return maxCell, maxDistance;
}
