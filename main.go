package main

import (
	"fmt"
	"mazes/model"
)

func main() {

	x := model.NewCell(0, 1);
	y := model.NewCell(0, 2);

	x.Link(y, true);
	fmt.Println(x);
}
