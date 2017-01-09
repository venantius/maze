package model

import (
	"testing"
	_ "fmt"
)

func TestNewGrid(t *testing.T) {
	x := NewGrid(2, 3);
	if x.rows != 2 {
		t.Fail()
	}
	if x.columns != 3 {
		t.Fail()
	}
}
