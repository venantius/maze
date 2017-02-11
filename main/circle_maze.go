package main

import (
	"maze/model"
	"fmt"
	_ "maze/generator"
	"maze/generator"
)

func main() {
	grid := model.NewPolarGrid(15);
	generator.Wilsons(grid);
	/*
	c1 := grid.GetCell(1, 0);
	c2 := grid.GetCell(1, 1);
	c1.Link(c2, true);

	c3 := grid.GetCell(0, 0);
	c1.Link(c3, true);
	*/

	filename := "circle_maze.png";
	grid.ToPNG(filename, 12);
	fmt.Printf("Saved to %v\n", filename);
}
