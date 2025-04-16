package main

import (
	"testing"
)

// Test the blinker oscillator pattern
func TestStep_Blinker(t *testing.T) {
	// Vertical line (blinker)
	input := map[Coord]bool{
		{1, 0}: true,
		{1, 1}: true,
		{1, 2}: true,
	}

	// After one step, it should be a horizontal line
	expected := map[Coord]bool{
		{0, 1}: true,
		{1, 1}: true,
		{2, 1}: true,
	}

	result := step(input)
	if !equal(result, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}

	// After a second step, it should return to vertical
	result2 := step(result)
	if !equal(result2, input) {
		t.Errorf("Expected: %v, Got: %v", input, result2)
	}
}

// Helper function to compare two maps
func equal(a, b map[Coord]bool) bool {
	if len(a) != len(b) {
		return false
	}
	for key := range a {
		if !b[key] {
			return false
		}
	}
	return true
}
