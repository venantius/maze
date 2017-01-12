package model

type Distances struct {
	root	*cell
	cells	map[*cell]int
}

func NewDistances(root *cell) *Distances {
	c := make(map[*cell]int);
	c[root] = 0;
	return &Distances {
		root: root,
		cells: c,
	}
}

func (d *Distances) Cells () <-chan *cell {
	ch := make(chan *cell, 1);
	go func() {
		for c, _ := range(d.cells) {
			ch <- c
		}
		close(ch);
	} ();
	return ch;
}
