package cell

type GridCell struct {
	*baseCell

	links  map[*GridCell]bool
	north  Cell
	east   Cell
	south  Cell
	west   Cell
}

func NewGridCell(row int, column int) *GridCell {
	return &GridCell{
		&baseCell{
			row: row,
			column: column,
		},
		make(map[*GridCell]bool),
		nil,
		nil,
		nil,
		nil,
	}
}


// These can probably be refactored somehow?
func (c *GridCell) North() Cell {
	return c.north;
}

func (c *GridCell) SetNorth(cell Cell) {
	c.north = cell;
}

func (c *GridCell) East() Cell {
	return c.east;
}

func (c *GridCell) SetEast(cell Cell) {
	c.east = cell;
}

func (c *GridCell) South() Cell {
	return c.south;
}

func (c *GridCell) SetSouth(cell Cell) {
	c.south = cell;
}

func (c *GridCell) West() Cell {
	return c.west;
}

func (c *GridCell) SetWest(cell Cell) {
	c.west = cell;
}

func (c *GridCell) Link(other Cell, bidi bool) {
	c.links[other.(*GridCell)] = true;
	if (bidi == true) {
		other.Link(c, false);
	}
}

func (c *GridCell) Unlink(other Cell, bidi bool) {
	delete(c.links, other.(*GridCell))
	if (bidi == true) {
		other.Unlink(c, false);
	}
}

func (c *GridCell) Links() []Cell {
	var keys []Cell = make([]Cell, 0, len(c.links))
	for k := range c.links {
		keys = append(keys, k)
	}
	return keys
}

// Is this cell linked to the other cell?
func (c *GridCell) IsLinked(other Cell) bool {
	if other == nil {
		return false;
	}
	_, exists := c.links[other.(*GridCell)]
	return exists
}

// All non-nil neighboring cells, whether linked or not.
func (c *GridCell) Neighbors() []Cell {
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

