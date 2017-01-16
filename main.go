package main

import (
	_ "fmt"
	"maze/model"
	"maze/generator"
	"fmt"
)

// I find this stuff helpful for reasoning about field embedding in Go.
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

func TestBaseGrid() {
	g := model.NewBaseGrid(5, 5);

	generator.BinaryTree(g);
	// model.Sidewinder(g);

	fmt.Println(g);
	g.ToPNG("derp.png", 10);
}

func TestDjikstra () {
	g := model.NewDistanceGrid(5, 5);
	generator.BinaryTree(g);

	start := g.RandomCell();
	distances := start.Distances();

	g.SetDistances(distances);
	fmt.Println(g);

	g.SetDistances(distances.PathTo(g.GetCell(g.GetRows() - 1, 0)));
	fmt.Println(g);
}

func TestLongestPath() {
	g := model.NewDistanceGrid(5, 5);
	generator.BinaryTree(g);

	start := g.GetCell(0, 0);

	distances := start.Distances();
	new_start, _ := distances.Max();

	new_distances := new_start.Distances();
	goal, _ := new_distances.Max();

	g.SetDistances(new_distances.PathTo(goal));
	fmt.Println(g);
	g.ToPNG("longest.png", 10);
}

func TestColoring() {
	g := model.NewColoredGrid(75, 75);
	generator.Sidewinder(g);

	start := g.GetCell(g.GetRows() / 2, g.GetColumns() / 2);

	g.SetDistances(start.Distances());

	filename := "colorized.png"
	g.ToPNG(filename, 10)
}

func TestAldousBroder() {
	g := model.NewBaseGrid(20, 20);
	generator.AldousBroder(g);

	filename := "aldous_broder.png"
	g.ToPNG(filename, 10);
}

func TestAldousBroderColored() {
	for i := 0; i < 6; i++ {
		g := model.NewColoredGrid(20, 20);
		generator.AldousBroder(g);

		middle := g.GetCell(g.GetRows() / 2, g.GetColumns() / 2);
		g.SetDistances(middle.Distances());

		filename := fmt.Sprintf("aldous_broder_%v.png", i);
		g.ToPNG(filename, 10);
	}


}

func main() {
	TestAldousBroderColored();
}
