package generator

import (
	"testing"
	"maze/model"
)

func TestBinaryTree_On(t *testing.T) {
	g := model.NewBaseGrid(4, 4);
	BinaryTree(g);
}
