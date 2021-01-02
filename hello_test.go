package main

import "testing"

func TestAddition(t *testing.T) {
	if addition(2, 2) != 4 {
		t.Error("Expected 2+2 to be equal to 4")
	}
}
