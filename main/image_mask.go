package main

import (
	"os"
	"log"
	"maze/model"
	"maze/generator"
	"fmt"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please specify a text file to use as a template.")
		os.Exit(1);
	}
	mask := model.MaskFromPNG(os.Args[1]);
	grid := model.NewMaskedGrid(mask);
	generator.RecursiveBacktracker(grid);

	filename := "masked.png"
	grid.ToPNG(filename, 10);
	fmt.Printf("Saved image to %v", filename);
}
