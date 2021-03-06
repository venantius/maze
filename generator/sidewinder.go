package generator

import (
	"maze/model"
	"maze/util"
	"maze/model/cell"
)

func max(x int, y int) int {
	if x > y {
		return x;
	}
	return y;
}

func Sidewinder(g model.Grid) {
	for row := range(g.RowIter()) {
		run := make([]*cell.GridCell, 0, 4);

		for _, c := range row {
			c := c.(*cell.GridCell) // lazy
			run = append(run, c);

			var at_eastern_boundary bool = c.East() == nil;
			var at_northern_boundary bool = c.North() == nil;

			should_close_out := at_eastern_boundary || (!at_northern_boundary && util.RANDOM.Intn(2) == 0)

			if should_close_out {
				member := run[util.RANDOM.Intn(len(run))];
				if member.North() != nil {
					member.Link(member.North(), true);
				}
				run = make([]*cell.GridCell, 0, max(g.Rows(), g.Columns()));
			} else {
				c.Link(c.East(), true);
			}
		}
	}
}
