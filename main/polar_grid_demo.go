package main

import (
	"maze/model"
	"fmt"
)

func main() {
	grid := model.NewPolarGrid(8);

	filename := "polar.png"
	grid.ToPNG(filename, 10);
	fmt.Printf("saved to %v\n", filename)
}
