package main

import "testing"

// EveOrOdd checks if a number is even or odd
func TestEvenOrOdd(t *testing.T){
	result := EvenOrOdd(10)
	if result != "even" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
