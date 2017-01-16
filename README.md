# Maze

This repository includes code for generating and solving mazes. The algorithms
included are taken from [Mazes for Programmers](https://pragprog.com/book/jbmaze/mazes-for-programmers),
and have been translated here from their original Ruby to Go.

## Examples

### Maze Generation

#### Grids

You can generate a grid maze like so:

```golang
// first, we create a 5x5 grid.
g := model.NewGrid(5, 5);

// Next, we apply the Sidewinder maze generation algorithm to it
model.Sidewinder(g);

// Now, we print it as ASCII
fmt.Println(g);
```

This will generate a random maze with Sidewinder that looks like the following:

```
+---+---+---+---+---+
|                   |
+   +   +   +---+   +
|   |   |   |       |
+---+---+   +   +   +
|           |   |   |
+---+---+   +---+   +
|           |       |
+   +---+   +   +   +
|   |       |   |   |
+---+---+---+---+---+
```

