package main

import (
	"testing"
)

// Test methods start with `Test`
func TestSum(t *testing.T) {
	got := Sum(1, 2)
	want := 3
	if got != want {
		t.Errorf("Sum(1, 2) == %d, want %d", got, want)
	}
}