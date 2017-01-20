package generator

import (
	"maze/model"
	"fmt"
)

func HuntAndKill(g model.Grid) {
	current := g.RandomCell();

	for current != nil {
		unvisited_neighbors := current.Neighbors();
		fmt.Println(unvisited_neighbors);
	}

}
