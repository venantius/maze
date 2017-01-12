package main

import (
	_ "fmt"
	_ "mazes/model"
	"mazes/model"
	"fmt"
)

type A struct {
	foo int
}

func (a *A) blah() {
	println("Foo!");
}

func (a *A) boop() {
	println("Zoom!");
}

type B struct {
	*A
	bar int
}

func (b *B) blah() {
	println("Bar!");
}


func djikstra_test () {
	g := model.NewDistanceGrid(5, 5);
	model.BinaryTree(g);

	start := g.GetCell(0, 0);
	distances := start.Distances();

	g.SetDistances(distances);
	fmt.Println(g);
}

func main() {

	djikstra_test();
	// Here's where our normal maze stuff begins.
	g := model.NewBaseGrid(5, 5);

	model.BinaryTree(g);
	// model.Sidewinder(g);

	fmt.Println(g);
	g.ToPNG(10);

}
