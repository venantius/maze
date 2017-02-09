package model

import (
	"strconv"
	"maze/util"
)

type Cell interface {
	Link(Cell, bool)
	Unlink(Cell, bool)

	Links() []Cell

	IsLinked(Cell) bool

	// Getters & Setters
	Row() int
	Col() int

	// BaseCell only, really
	North() Cell
	SetNorth(Cell)
	East() Cell
	SetEast(Cell)
	South() Cell
	SetSouth(Cell)
	West() Cell
	SetWest(Cell)

	Neighbors() []Cell
	Distances() *Distances
}

type BaseCell struct {
	column int
	row    int

	links  map[Cell]bool

	north  Cell
	east   Cell
	south  Cell
	west   Cell
}

func NewBaseCell(row int, column int) *BaseCell {
	return &BaseCell{
		row: row,
		column: column,
		links: make(map[Cell]bool),
	}
}

func (c *BaseCell) Row() int {
	return c.row;
}

func (c *BaseCell) Col() int {
	return c.column;
}

// These can probably be refactored somehow?
func (c *BaseCell) North() Cell {
	return c.north;
}

func (c *BaseCell) SetNorth(cell Cell) {
	c.north = cell;
}

func (c *BaseCell) East() Cell {
	return c.east;
}

func (c *BaseCell) SetEast(cell Cell) {
	c.east = cell;
}

func (c *BaseCell) South() Cell {
	return c.south;
}

func (c *BaseCell) SetSouth(cell Cell) {
	c.south = cell;
}

func (c *BaseCell) West() Cell {
	return c.west;
}

func (c *BaseCell) SetWest(cell Cell) {
	c.west = cell;
}

func (c *BaseCell) Link(other Cell, bidi bool) {
	c.links[other] = true;
	if (bidi == true) {
		other.Link(c, false);
	}
}

func (c *BaseCell) Unlink(other Cell, bidi bool) {
	delete(c.links, other)
	if (bidi == true) {
		other.Unlink(c, false);
	}
}

func (c *BaseCell) Links() []Cell {
	var keys []Cell = make([]Cell, 0, len(c.links))
	for k := range c.links {
		keys = append(keys, k)
	}
	return keys
}

// Is this cell linked to the other cell?
func (c *BaseCell) IsLinked(other Cell) bool {
	_, exists := c.links[other]
	return exists
}

// All non-nil neighboring cells, whether linked or not.
func (c *BaseCell) Neighbors() []Cell {
	output := make([]Cell, 0, 4);
	if c.North() != nil {
		output = append(output, c.North());
	}
	if c.East() != nil {
		output = append(output, c.East());
	}
	if c.South() != nil {
		output = append(output, c.South());
	}
	if c.West() != nil {
		output = append(output, c.West());
	}
	return output;
}

// String representation.
func (c *BaseCell) String() string {
	output := "{:row " + strconv.Itoa(c.row);
	output += " :column " + strconv.Itoa(c.column);
	output += "}"
	return output
}



// Part 1 of an implementation of Djikstra's graph search algorithm as applied to mazes.
func (c *BaseCell) Distances() *Distances {
	distances := NewDistances(c);
	frontier := []Cell{c};

	for SliceHasAny(frontier) {
		new_frontier := []Cell{};

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

// Utility functions for working with slices of *Cell structs

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
