package main

import (
	"fmt"
	"mazes/model"
)

func main() {

	// b := model.NewCell(0, 1);

	g := model.NewGrid(2, 2);
	fmt.Printf("%v", g);

	for row := range g.Cells() {
		fmt.Println(row);
	}
}
