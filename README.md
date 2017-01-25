# Overview

This repository includes code for generating and solving mazes. The algorithms
included are taken from [Mazes for Programmers](https://pragprog.com/book/jbmaze/mazes-for-programmers),
and have been translated here from their original Ruby to Go.

This README covers the basics of how to work with this library and what sort of output you can expect from it.

## Grids

Grids are the basic "objects" of mazes. The general usage pattern is to
create a grid and then to apply a maze generation algorithm to it. From there, you can run additional algorithms, for instance to determine the shortest or longest path from one point in the maze to another.

### BaseGrid

The most basic type of grid in this library is a `BaseGrid`. You can see how it works as follows:

```golang
import (
    "maze.generator"
    "maze.model"
)

func main() {
    // first, we create a 5x5 BaseGrid.
    g := model.NewBaseGrid(5, 5);

    // Next, we apply the Sidewinder maze generation algorithm to it
    generator.Sidewinder(g);

    // Now, we print it as ASCII
    fmt.Println(g);
}
```

This will print something like the following to stdout:

```
+---+---+---+---+---+
|                   |
+   +---+   +   +---+
|       |   |       |
+   +   +---+   +---+
|   |       |       |
+   +   +---+   +---+
|   |   |           |
+---+   +   +---+---+
|       |           |
+---+---+---+---+---+
```

You can also generate a PNG image of the same grid using `ToPNG`:

```golang
filename := "sidewinder.png";
g.ToPNG(filename, 10);
```

This will generate an image that looks like the following:

![](/doc/images/sidewinder.png)

### DistanceGrid

If you want to see how far a given point is from another, you can use a
`DistanceGrid`. In practice, `DistanceGrid` by itself isn't that interesting,
but is a good foundation for gradient coloring later on.

The following example shows how you can use a `DistanceGrid` to determine
the shortest path between any two points in a maze, using Djikstra's algorithm:

```go
func main() {
    g := model.NewDistanceGrid(5, 5)

    // Apply the Binary Tree maze generation algorithm
    generator.BinaryTree(g)

    // Pick a random starting cell, or any cell you want.
    start := g.RandomCell();

    // Calculate the shortest path between the starting cell and all other cells.
    distances := start.Distances();
    g.SetDistances(distances);

    // Print to stdout.
    fmt.Println(g)
}
```

Note that distances will be measured in base36, meaning the 10th cell away from
the starting point will be assigned an 'a', the 11th cell will be a 'b', etc.

```
+---+---+---+---+---+
| 3   2   3   4   5 |
+   +   +   +---+   +
| 4 | 1 | 4 | 7   6 |
+   +   +   +   +   +
| 5 | 0 | 5 | 8 | 7 |
+   +---+   +   +   +
| 6 | 7   6 | 9 | 8 |
+   +   +---+---+   +
| 7 | 8 | b   a   9 |
+---+---+---+---+---+
```

You can also extend our example to directly print the shortest path between
two points, like so :

```go
// Print the shortest distance between our random starting cell and the cell
// in the bottom-left
g.SetDistances(distances.PathTo(g.GetCell(g.GetRows() - 1, 0)));
fmt.Println(g);
```

This will print something like the following:

```
+---+---+---+---+---+
| 3   2             |
+   +   +   +---+   +
| 4 | 1 |   |       |
+   +   +   +   +   +
| 5 | 0 |   |   |   |
+   +---+   +   +   +
| 6 |       |   |   |
+   +   +---+---+   +
| 7 |   |           |
+---+---+---+---+---+
```

### ColoredGrid

If you want to generate colored PNG images of your mazes, use the `ColoredGrid` struct instead of `BaseGrid`, like so
