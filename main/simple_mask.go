package main

import (
	"maze/model"
	"maze/generator"
	"fmt"
)

func main() {
	mask := model.NewMask(5, 5);

	mask.SetCell(0, 0, false);
	mask.SetCell(2, 2, false);
	mask.SetCell(4, 4, false);

	grid := model.NewMaskedGrid(mask);
	generator.RecursiveBacktracker(grid);

	fmt.Println(grid);
}
