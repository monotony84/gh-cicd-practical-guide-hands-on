package main

import "testing"

func testEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("expected: Even, actual: %s", result)
	}
}