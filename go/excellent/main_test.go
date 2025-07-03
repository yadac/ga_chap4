package main

import "testing"

// EveOrOdd checks if a number is even or odd
func TestEvenOrOdd(t *testing.T){
	result := EveOrOdd(10)
	if result != "even" {
		t.Error("expected: even, actual: %s", result)
	}
}
