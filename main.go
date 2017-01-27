package main

import (
	_ "fmt"
	"maze/model"
	"maze/generator"
	"fmt"
)

func TestAlgorithm(algo generator.MazeGen) {
	g := model.NewBaseGrid(30, 30);
	algo(g);

	filename := "maze.png";
	g.ToPNG(filename, 10);
}

func FindLongestPath(g model.IDistanceGrid) {
	random := g.RandomCell();
	distances := random.Distances();
	new_start, _ := distances.Max(); // furthest cell from our random cell

	new_distances := new_start.Distances();

	g.SetDistances(new_distances);
}

func TestAlgorithmColored(algo generator.MazeGen) {
	g := model.NewColoredGrid(30, 30);
	algo(g);

	FindLongestPath(g);

	filename := "colored_maze.png";
	g.ToPNG(filename, 10);
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

	g.SetDistances(new_distances);
	fmt.Println(g);
	g.ToPNG("longest.png", 10);
}

func TestBinaryTree() {
	TestAlgorithm(generator.BinaryTree);
}

func TestSidewinder() {
	TestAlgorithm(generator.Sidewinder);
}

func TestAldousBroder() {
	TestAlgorithm(generator.AldousBroder);
}

func TestAldousBroderColored() {
	TestAlgorithmColored(generator.AldousBroder);
}

func TestWilsons() {
	TestAlgorithm(generator.Wilsons);
}

func TestWilsonsColored() {
	TestAlgorithmColored(generator.Wilsons);
}

func TestHuntAndKill() {
	TestAlgorithm(generator.HuntAndKill);
}

func TestHuntAndKillColored() {
	TestAlgorithmColored(generator.HuntAndKill);
}

func TestRecursiveBacktracker() {
	TestAlgorithm(generator.RecursiveBacktracker);
}

func TestRecursiveBacktrackerColored() {
	TestAlgorithmColored(generator.RecursiveBacktracker);
}

func main() {
	TestRecursiveBacktrackerColored();
}

