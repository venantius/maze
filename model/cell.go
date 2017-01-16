package model

import (
	"strconv"
)

type Cell struct {
	column int
	row    int

	links  map[*Cell]bool

	North  *Cell
	East   *Cell
	South  *Cell
	West   *Cell
}

func NewCell(row int, column int) *Cell {
	return & Cell{
		row: row,
		column: column,
		links: make(map[*Cell]bool),
	}
}

func (c *Cell) Link(other *Cell, bidi bool) {
	c.links[other] = true;
	if (bidi == true) {
		other.Link(c, false);
	}
}

func (c *Cell) Unlink(other *Cell, bidi bool) {
	delete(c.links, other)
	if (bidi == true) {
		other.Unlink(c, false);
	}
}

func (c *Cell) Links() []*Cell {
	var keys []*Cell = make([]*Cell, 0, len(c.links))
	for k := range c.links {
		keys = append(keys, k)
	}
	return keys
}

// Is this cell linked to the other cell?
func (c *Cell) IsLinked(other *Cell) bool {
	_, exists := c.links[other]
	return exists
}

// All non-nil neighboring cells, whether linked or not.
func (c *Cell) Neighbors() []*Cell {
	output := make([]*Cell, 0, 4);
	if c.North != nil {
		output = append(output, c.North);
	}
	if c.East != nil {
		output = append(output, c.East);
	}
	if c.South != nil {
		output = append(output, c.South);
	}
	if c.West != nil {
		output = append(output, c.West);
	}
	return output;
}

// String representation.
func (c *Cell) String() string {
	output := "{:row " + strconv.Itoa(c.row);
	output += " :column " + strconv.Itoa(c.column);
	output += "}"
	return output
}

// Check to see if any of the *Cells in this slice are not nil. Akin to Ruby's `any?`
func hasAny (cells []*Cell) bool {
	for _, c := range cells  {
		if c != nil {
			return true;
		}
	}
	return false
}

// Part 1 of an implementation of Djikstra's graph search algorithm as applied to mazes.
func (c *Cell) Distances() *Distances {
	distances := NewDistances(c);
	frontier := []*Cell{c};

	for hasAny(frontier) {
		new_frontier := []*Cell{};

		for _, cell := range(frontier) {
			for _, linked := range(cell.Links()) {
				_, ok := distances.cells[linked]
				if ok {
					continue
				}
				distances.cells[linked] = distances.cells[cell] + 1;
				new_frontier = append(new_frontier, linked);
			}
		}
		frontier = new_frontier;
	}
	return distances;
}
