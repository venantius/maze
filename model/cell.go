package model

import (
	_ "fmt"
	"strconv"
)

type cell struct {
	column int
	row    int

	links  map[*cell]bool

	north  *cell
	east   *cell
	south  *cell
	west   *cell
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

func (c *cell) Unlink(other *cell, bidi bool) {
	delete(c.links, other)
	if (bidi == true) {
		other.Unlink(c, false);
	}
}

func (c *cell) Links() []*cell {
	var keys []*cell = make([]*cell, 0, len(c.links))
	for k := range c.links {
		keys = append(keys, k)
	}
	return keys
}

// Is this cell linked to the other cell?
func (c *cell) IsLinked(other *cell) bool {
	_, exists := c.links[other]
	return exists
}

func (c *cell) String() string {
	output := "{:row " + strconv.Itoa(c.row);
	output += " :column " + strconv.Itoa(c.column);
	output += "}"
	return output
}


