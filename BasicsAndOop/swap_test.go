package main

import (
	"testing"
)

func TestSwapReturn(t *testing.T) {
	var a, b = 10, 11
	a, b = swapReturn(10, 11)
	if a != 11 || b != 10 {
		t.Error("a should be 11 and b should be 10")
	}
}

func TestSwapPointer(t *testing.T) {
	var a, b = 10, 11
	swapPointer(&a, &b)
	if a != 11 || b != 10 {
		t.Error("a should be 11 and b should be 10")
	}
}
