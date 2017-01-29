package main

import (
	"fmt"
	"maze/model"
	"maze/generator"
)

func main() {
	g := model.NewBaseGrid(5, 5);
	g.GetCell(0, 0).East.West = nil;
	g.GetCell(0, 0).South.North = nil;

	g.GetCell(4, 4).West.East = nil;
	g.GetCell(4, 4).North.South = nil;

	generator.RecursiveBacktracker(g);
	fmt.Println(g)
}
