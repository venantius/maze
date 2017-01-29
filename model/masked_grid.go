package model

/*
Chapter 6
 */

type maskedGrid struct {
	*baseGrid

	mask *mask
}

func NewMaskedGrid(mask *mask) *maskedGrid {
	m := &maskedGrid{
		&baseGrid{
			rows: mask.rows,
			columns: mask.columns,
		},
		mask,
	}
	m.grid = m.prepareGrid();
	m.configureCells();
	return m;
}

func (m *maskedGrid) prepareGrid() [][]*Cell {
	grid := make([][]*Cell, m.rows)
	for i, _ := range(grid) {
		column := make([]*Cell, m.columns)
		grid[i] = column
		for j, _ := range(column) {
			if m.mask.bits[i][j] == true {
				grid[i][j] = NewCell(i, j);
			}
		}
	}
	return grid;
}

func (m *maskedGrid) RandomCell() *Cell {
	row, col := m.mask.RandomCell();
	return m.grid[row][col];
}

func (m *maskedGrid) Size() int {
	return m.mask.Count();
}
