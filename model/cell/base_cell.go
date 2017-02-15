package cell

import "strconv"

// An extremely minimal basic struct for tracking rows and columns.
// Nothing special happening here.
type baseCell struct {
	column int
	row    int
}

func (b *baseCell) Row() int {
	return b.row;
}

func (b *baseCell) Column() int {
	return b.column;
}

func (g *GridCell) Distances() *Distances {
	return calculateDistances(g);
}

func (c *baseCell) Neighbors() []Cell {
	panic("Not yet implemented");
}

// String representation.
func (c *baseCell) String() string {
	output := "{:row " + strconv.Itoa(c.row);
	output += " :column " + strconv.Itoa(c.column);
	output += "}"
	return output
}
