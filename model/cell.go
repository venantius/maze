package model

type cell struct {
	column int
	row    int

	links  map[*cell]bool
}

func NewCell(row int, column int) *cell {
	return & cell{
		row: row,
		column: column,
		links: make(map[*cell]bool),
	}
}

func (c *cell) Link(other *cell, bidi bool) {
	c.links[other] = true;
	if (bidi == true) {
		other.Link(c, false);
	}
}

func (c *cell) String() string {

}


