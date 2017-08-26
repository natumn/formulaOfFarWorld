package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	calcStr := "15*5"
	answer := calc(calcStr)
	if answer != 75 {
		t.Fatal("miss calculator.")
	}

	t.Log(calcStr, "=", answer)
}
