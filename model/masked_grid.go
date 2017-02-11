package model

import "maze/model/cell"

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

func (m *maskedGrid) prepareGrid() [][]cell.Cell {
	grid := make([][]cell.Cell, m.rows)
	for i, _ := range(grid) {
		column := make([]cell.Cell, m.columns)
		grid[i] = column
		for j, _ := range(column) {
			if m.mask.bits[i][j] == true {
				grid[i][j] = cell.NewGridCell(i, j);
			}
		}
	}
	return grid;
}

func (m *maskedGrid) RandomCell() cell.Cell {
	row, col := m.mask.RandomCell();
	return m.grid[row][col];
}

func (m *maskedGrid) Size() int {
	return m.mask.Count();
}
