package model

import (
	_ "fmt"
)

func BinaryTree (g grid) {
	for c := range g.Cells() {
		neighbors := make([]*cell, 0, 2);

		if c.north != nil {
			neighbors = append(neighbors, c.north);
		}
		if c.east != nil {
			neighbors = append(neighbors, c.east);
		}

		if len(neighbors) != 0 {
			index := RANDOM.Intn(len(neighbors));
			neighbor := neighbors[index];
			c.Link(neighbor, true);
		}
	}
}


