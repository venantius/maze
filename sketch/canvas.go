package sketch

import (
	"image/color"
	"image/draw"
	"math"
	"image"
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

type tier struct {
	xMin int
	xMax int
}

func updateTier(x int, y int, tiers map[int]tier) {
	currentTier, exists := tiers[y];
	if !exists {
		currentTier = tier{x, x};
	}
	if currentTier.xMax <= x {
		currentTier.xMax = x;
	}
	if currentTier.xMin >= x {
		currentTier.xMin = x;
	}
	tiers[y] = currentTier;
}

func setTiers(p1 image.Point, p2 image.Point, tiers map[int]tier) {
	var xMin, xMax, yMin, yMax int;

	if p1.X < p2.X {
		xMin, xMax = p1.X, p2.X;
	} else {
		xMin, xMax = p2.X, p1.X;
	}

	if p1.Y < p2.Y {
		yMin, yMax = p1.Y, p2.Y;
	} else {
		yMin, yMax = p2.Y, p1.Y;
	}

	var slope float64 = float64(p2.Y - p1.Y) / float64(p2.X - p1.X);
	var yIntercept float64 = float64(p1.Y) - slope * float64(p1.X);

	if math.IsNaN(slope) {
		// 0 divided by 0; draw nothing.
		panic("You're not allowed to add points that have the same location and call it a quadrilateral")
	} else if math.IsInf(slope, 0) {
		// any non-zero divided by 0, vertical line
		x := xMin;
		for y := yMin; y <= yMax; y++ {
			updateTier(x, y, tiers);
		}
	} else {
		// a normal slope
		// if the slope is <= 1, iterate over x. Otherwise, iterate over y.
		// this prevents us from skipping pixels if the slope is too steep
		if math.Abs(slope) <= 1 {
			for x := xMin; x <= xMax; x++ {
				y := int(float64(x) * slope + yIntercept);
				updateTier(x, y, tiers);
			}
		} else {
			for y := yMin; y <= yMax; y++ {
				x := int((float64(y) - yIntercept) / slope);
				updateTier(x, y, tiers);
			}
		}
	}
}

// Draw a quadrilateral, pixel by pixel. No anti-aliasing.
func DrawQuadrilateral(points []image.Point, img draw.Image, col color.Color, wall color.Color) {

	// a map of y-position to a range of xmin, xmaxes that need to be colored.
	tiers := make(map[int]tier);

	// populate tiers
	for i := 0; i < len(points); i++ {
		var p2 image.Point;
		p1 := points[i];
		if i == len(points) - 1 {
			p2 = points[0];
		} else {
			p2 = points[i + 1];
		}
		setTiers(p1, p2, tiers);
	}

	// fill tiers
	for y, t := range(tiers) {
		for x := t.xMin; x < t.xMax; x++ {
			// Don't paint over walls.
			if img.At(x, y) != wall {
				img.Set(x, y, col);
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