package main

import (
	_ "maze/model"
	"maze/generator"
	"fmt"
	"maze/model"
)

func main() {
	// I probably could have done this with reflection but didn't feel like it.
	var algorithms map[string]generator.MazeGen = map[string]generator.MazeGen{
		"BinaryTree":    generator.BinaryTree,
		"Sidewinder":    generator.Sidewinder,
		"Aldous Broder": generator.AldousBroder,
		"Wilson's": 	 generator.Wilsons,
		"Hunt and Kill": generator.HuntAndKill,
	};

	var tries int = 100;
	var size int = 20;

	var averages map[string]int = map[string]int{};

	for name, algorithm := range(algorithms) {
		fmt.Printf("running %v\n", name);

		deadend_counts := []int{};
		for i := 0; i < tries; i++ {
			var grid model.Grid = model.NewBaseGrid(size, size);
			algorithm(grid);

			deadend_counts = append(deadend_counts, len(grid.Deadends()));
		}

		total_deadends := 0;
		for _, count := range(deadend_counts) {
			total_deadends += count;
		}

		averages[name] = total_deadends / len(deadend_counts);
	}

	total_cells := size * size;

	fmt.Printf("\nAverage dead-ends per %vx%v maze (%v cells):\n\n", size, size, total_cells);


	// TODO: Sort this first by averages.
	for name, _ := range(algorithms) {
		percentage := averages[name] * 100.0 / (size * size);
		fmt.Printf("%v : %v/%v (%v)\n", name, averages[name], total_cells, percentage);
	}
}
