package main

import "testing"

func TestEvenOrOdd(t *testing.T){
	result := EveOrOdd(10)
	if result != "even" {
		t.Error("expected: even, actual: %s", result)
	}
}
