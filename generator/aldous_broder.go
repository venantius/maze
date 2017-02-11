package generator

import (
	"maze/model"
	"maze/util"
	"fmt"
	"maze/model/cell"
)

func AldousBroder (g model.Grid) {
	c := g.RandomCell();
	fmt.Printf("%T\n", c);

	var unvisited int = g.Size() - 1;

	for unvisited > 0 {
		var neighbor cell.Cell = c.Neighbors()[util.RANDOM.Intn(len(c.Neighbors()))];
		fmt.Printf("%T\n", c);

		if len(neighbor.Links()) == 0 {
			c.Link(neighbor, true);
			unvisited -= 1
		}

		c = neighbor;
	}
}