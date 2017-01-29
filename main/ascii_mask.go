package main

import (
	"maze/model"
	"os"
	"log"
	"maze/generator"
	"fmt"
)

// Sample usage: ./ascii_masc main/mask.txt
func main() {
	if len(os.Args) == 1 {
		log.Fatal("Please specify a text file to use as a template.")
		os.Exit(1);
	}
	mask := model.MaskFromText(os.Args[1])
	grid := model.NewMaskedGrid(mask);

	generator.RecursiveBacktracker(grid);

	filename := "masked.png";
	grid.ToPNG(filename, 10);
	fmt.Printf("saved image to %v", filename);
}
