package model

import _ "fmt"

func max(x int, y int) int {
	if x > y {
		return x;
	}
	return y;
}

func Sidewinder(g *baseGrid) {
	for row := range(g.RowIter()) {
		run := make([]*cell, 0, 4);

		for _, c := range row {
			run = append(run, c);

			var at_eastern_boundary bool = c.east == nil;
			var at_northern_boundary bool = c.north == nil;

			should_close_out := at_eastern_boundary || (!at_northern_boundary && RANDOM.Intn(2) == 0)

			if should_close_out {
				member := run[RANDOM.Intn(len(run))];
				if member.north != nil {
					member.Link(member.north, true);
				}
				run = make([]*cell, 0, max(g.rows, g.columns));
			} else {
				c.Link(c.east, true);
			}
		}
	}
}
