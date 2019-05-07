package main

import (
	"testing"
)

func TestShouldAssertTrianguleArea(t *testing.T) {
	triangle := triangle{base: 2, height: 3}
	if triangle.getArea() != 3 {
		t.Errorf("Test expected 3")
	}
}

func TestShouldAssertSquareArea(t *testing.T) {
	square := square{sideLength: 2}
	if square.getArea() != 4 {
		t.Errorf("Test expected 4")
	}
}

func TestShouldAssertExecuteFunction(t *testing.T) {
	printArea(square{sideLength: 2})
}
