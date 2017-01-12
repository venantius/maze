package model

import (
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

func hasAny (cells []*cell) bool {
	for _, c := range cells  {
		if c != nil {
			return true;
		}
	}
	return false
}

// Part 1 of an implementation of Djikstra's graph search algorithm as applied to mazes.
func (c *cell) Distances() *Distances {
	distances := NewDistances(c);
	frontier := []*cell{c};

	for hasAny(frontier) {
		new_frontier := []*cell{};

		for _, celllllll := range(frontier) {
			for _, linked := range(celllllll.Links()) {
				_, ok := distances.cells[linked]
				if ok {
					continue
				}
				distances.cells[linked] = distances.cells[celllllll] + 1;
				new_frontier = append(new_frontier, linked);
			}
		}
		frontier = new_frontier;
	}
	return distances;
}
