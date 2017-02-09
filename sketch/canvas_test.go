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
