package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	numSlice := []int{1}
	resString1 := evenOrOdd(numSlice)
	expectedString1 := "1 is odd\n"
	if resString1 != expectedString1 {
		t.Errorf("expected %v, got %v", expectedString1, resString1)
	}

	numSlice2 := []int{2}
	resString2 := evenOrOdd(numSlice2)
	expectedString2 := "2 is even\n"
	if resString2 != expectedString2 {
		t.Errorf("expected %v, got %v", expectedString2, resString2)
	}

	numSlice3 := []int{3, 4}
	resString3 := evenOrOdd(numSlice3)
	expectedString3 := "3 is odd\n4 is even\n"
	if resString3 != expectedString3 {
		t.Errorf("expected %v, got %v", expectedString3, resString3)
	}

	numSlice4 := []int{}
	resString4 := evenOrOdd(numSlice4)
	expectedString4 := ""
	if resString4 != expectedString4 {
		t.Errorf("expected %v, got %v", expectedString4, resString4)
	}

}
