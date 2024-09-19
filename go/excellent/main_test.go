package main

import "testing"

func TestEventOrOdd(t *testing.T) {
	result := EventOrOdd(8)
	if result != "event" {
		t.Errorf("expected: even, actual: %s", result)
	}
}
