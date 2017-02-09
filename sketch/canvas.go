package sketch

import (
	"image/color"
	"image/draw"
	"math"
)

// Draw a line between two points
func DrawLine(x1 int, y1 int, x2 int, y2 int, img draw.Image, color color.Color) {
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

	var slope float64 = float64(y2 - y1) / float64(x2 - x1);
	var yIntercept float64 = float64(y1) - slope * float64(x1);

	if math.IsNaN(slope) {
		// 0 divided by 0; draw nothing.
	} else if math.IsInf(slope, 0) {
		// any non-zero divided by 0, vertical line
		for y := ymin; y <= ymax; y++ {
			img.Set(xmin, y, color);
		}
	} else {
		// a normal slope
		// if the slope is <= 1, iterate over x. Otherwise, iterate over y.
		// this prevents us from skipping pixels if the slope is too steep
		if math.Abs(slope) <= 1 {
			for x := xmin; x <= xmax; x++ {
				y := int(float64(x) * slope + yIntercept);
				img.Set(x, y, color);
			}
		} else {
			for y := ymin; y <= ymax; y++ {
				x := int((float64(y) - yIntercept) / slope);
				img.Set(x, y, color);
			}
		}
	}
}

func DrawRectangle(x1 int, y1 int, x2 int, y2 int, img draw.Image, color color.Color) {
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