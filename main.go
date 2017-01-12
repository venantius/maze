package main

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	_ "fmt"
	_ "mazes/model"
	"image/color"
	_ "mazes/model"
	"mazes/model"
	"fmt"
)


func testdraw2dimg() {
	dest := image.NewRGBA(image.Rect(0, 0, 50, 50));
	gc := draw2dimg.NewGraphicContext(dest);

	// Set some properties
	gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xff});
	gc.SetLineWidth(1)

	// Draw a closed shape
	gc.MoveTo(10, 10) // should always be called first for a new path
	gc.LineTo(10, 20)

	gc.Close();
	gc.FillStroke();

	draw2dimg.SaveToPngFile("hello.png", dest);
}

func main() {
	// testdraw2dimg();


	// Here's where our normal maze stuff begins.
	g := model.NewGrid(5, 5);

	model.BinaryTree(g);
	// model.Sidewinder(g);

	fmt.Println(g);
	g.ToPNG(10);

}
