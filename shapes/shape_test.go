package main

import "testing"

func TestNewSquare(t *testing.T) {
	s := square{sideLength: 2}

	if s.getArea() != 4 {
		t.Errorf("Expected square's area to be 4, got %v\n", s.getArea())
	}
}

func TestNewTriangle(t *testing.T) {
	tri := triangle{base: 4, height: 4}

	if tri.getArea() != 8 {
		t.Errorf("Expected triangle's area to be 8, got %v\n", tri.getArea())
	}
}
