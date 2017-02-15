package sketch

import (
	"testing"
	"image"
	"image/color"
	"github.com/llgcode/draw2d/draw2dimg"
)

func TestDrawLine(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100));

	// Draw horizontal line
	DrawLine(10, 10, 20, 10, img, color.Black);

	// Draw vertical line
	DrawLine(10, 10, 10, 20, img, color.Black);

	// Draw diagonal line from top left to bottom right
	DrawLine(10, 10, 20, 20, img, color.Black);

	// Draw diagonal line from the bottom left to the top right;
	DrawLine(10, 80, 20, 60, img, color.Black);
	draw2dimg.SaveToPngFile("testDrawLine.png", img);
}

func TestDrawQuadrilateral(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100));

	points := []image.Point{image.Point{10, 10}, image.Point{10, 30}, image.Point{30, 30}, image.Point{30, 10}}
	DrawQuadrilateral(points, img, color.RGBA{255, 0, 0, 255}, color.Black);

	points = []image.Point{image.Point{50, 10}, image.Point{60, 20}, image.Point{50, 30}, image.Point{40, 20}}
	DrawQuadrilateral(points, img, color.RGBA{255, 255, 0, 255}, color.Black);

	points = []image.Point{image.Point{10, 40}, image.Point{30, 50}, image.Point{20, 60}, image.Point{10, 55}}
	DrawQuadrilateral(points, img, color.RGBA{255, 0, 255, 255}, color.Black);
	draw2dimg.SaveToPngFile("testDrawQuadrilateral.png", img);
}