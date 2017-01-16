package generator

import (
	"maze/model"
	"maze/util"
)

func max(x int, y int) int {
	if x > y {
		return x;
	}
	return y;
}

func Sidewinder(g model.Grid) {
	for row := range(g.RowIter()) {
		run := make([]*model.Cell, 0, 4);

		for _, c := range row {
			run = append(run, c);

			var at_eastern_boundary bool = c.East == nil;
			var at_northern_boundary bool = c.North == nil;

			should_close_out := at_eastern_boundary || (!at_northern_boundary && util.RANDOM.Intn(2) == 0)

			if should_close_out {
				member := run[util.RANDOM.Intn(len(run))];
				if member.North != nil {
					member.Link(member.North, true);
				}
				run = make([]*model.Cell, 0, max(g.GetRows(), g.GetColumns()));
			} else {
				c.Link(c.East, true);
			}
		}
	}
}
