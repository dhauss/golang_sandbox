package main

import "testing"

func TestNewSquare(t *testing.T) {
	s := square{sideLength: 2}

	if s.getArea() != 4 {
		t.Errorf("Expected area of 4, got %v\n", s.getArea())
	}

}
