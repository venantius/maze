package sketch

import (
	"image"
	"image/color"
)

// Draw a line between two points
func DrawLine(x1 int, y1 int, x2 int, y2 int, img *image.RGBA, color color.Color) {
	var xmin, xmax, ymin, ymax int;

	if x1 < x2 {
		xmin, xmax = x1, x2;
	} else {
		xmin, xmax = x2, x1;
	}

	if y1 < y2 {
		ymin, ymax = y1, y2;
	} else {
		ymin, ymax = y2, y1;
	}

	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			img.Set(x, y, color);
		}
	}
}