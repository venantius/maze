# Maze

This repository includes code for generating and solving mazes. The algorithms
included are taken from [Mazes for Programmers](https://pragprog.com/book/jbmaze/mazes-for-programmers),
and have been translated here from their original Ruby to Go.

## Overview

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

### ColoredGrid

If you want to generate colored PNG images of your mazes, use the `ColoredGrid` struct instead of `BaseGrid`, like so (EXAMPLE TO FOLLOW)
