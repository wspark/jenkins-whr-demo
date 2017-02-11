package main

import (
	"testing"
)

func TestAdder(t *testing.T) {
	res := adder(1, 1)
	if res != 2 {
		t.Errorf("expected 2, got: %d", res)
	}
}
