package model

import "testing"

func TestBinaryTree_On(t *testing.T) {
	g := NewGrid(4, 4);
	bt := &BinaryTree{};
	bt.On(g);
}
